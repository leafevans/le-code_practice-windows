package interview

import (
	"context"
	"io"
	"log"
	"time"

	"ai-eino-interview-agent/api/handler/interview/mianshi"
	interviews "ai-eino-interview-agent/api/model/interviews"
	mianshiapi "ai-eino-interview-agent/api/model/mianshi"
	"ai-eino-interview-agent/api/response"
	"ai-eino-interview-agent/chatApp/agent_service/evaluation"
	"ai-eino-interview-agent/internal/middleware"
	interviewservice "ai-eino-interview-agent/internal/service/interviews"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// StartMianshiStream .
// @router /api/mianshi/stream/start [POST]
// 开始面试流程，返回一个 SSE 流，持续推送面试问题和反馈，直到面试结束
func StartMianshiStream(ctx context.Context, c *app.RequestContext) {
	// 解析请求参数
	var err error
	var req mianshiapi.MianshiStartInterviewRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 用户认证
	userID := middleware.GetUserID(c)
	if userID == 0 {
		// 未认证用户，返回 401 Unauthorized
		response.Unauthorized(ctx, c, "Unauthorized")
		return
	}

	// 调用服务层开始面试
	interviewService := interviewservice.NewInterviewService()
	// 创建面试记录DTO
	recordDTO := &interviews.InterviewRecordDTO{
		UserID:       int32(userID),
		Type:         req.Type,
		Difficulty:   req.Difficulty,
		Domain:       req.Domain,
		PositionName: req.PositionName,
		CompanyName:  req.CompanyName,
		Status:       "pending",
	}

	// 创建面试记录
	recordID, err := interviewService.CreateInterviewRecord(ctx, recordDTO)
	if err != nil {
		response.InternalServerError(ctx, c, "Failed to create interview record:"+err.Error())
		return
	}

	// 如果请求中包含简历ID，关联简历和面试记录
	var hasResume bool
	var resumeID int64
	if req.ResumeID != nil && *req.ResumeID > 0 {
		hasResume = true
		resumeID = *req.ResumeID
	}

	// 面试管理器，负责管理面试会话和状态
	manager := mianshi.GetGlobalSessionManager()
	// 这里可以根据需要传递更多的参数，比如职位信息、公司信息等
	var companyName, positionName string
	if req.CompanyName != nil {
		companyName = *req.CompanyName
	}
	if req.PositionName != nil {
		positionName = *req.PositionName
	}
	// 创建面试会话，记录用户ID、面试记录ID、简历ID等信息
	session := manager.CreateSessionWithDetails(userID, recordID, resumeID, hasResume, "", req.Type, req.Domain, req.Difficulty, companyName, positionName)

	// 启动面试流程
	cancel, cancelFunc := context.WithCancel(ctx)
	session.CancelFunc = cancelFunc

	// 启动面试流程的goroutine
	mianshi.SetupSSEResponse(c)

	// 创建一个管道，将面试问题和反馈写入管道，HTTP 响应从管道读取，实现 SSE 流式传输
	reader, writer := io.Pipe()
	c.SetBodyStream(reader, -1)

	// 在一个新的 goroutine 中运行面试流程，持续向管道写入数据，直到面试结束或被取消
	go func() {
		// 确保在面试流程结束时关闭管道写入端，通知 HTTP 响应流结束
		defer func(writer *io.PipeWriter) {
			err = writer.Close()
			if err != nil {
				log.Printf("Failed to close pipe writer: sessionID=%s, error=%v", session.SessionID, err)
				return
			}
		}(writer)
		// 面试流程开始，发送初始事件通知客户端面试已开始
		startTime := session.StartTime.UnixMilli()
		err := mianshi.SendSSEEvent(writer, map[string]any{
			"type":       "session_id",
			"session_id": session.SessionID,
			"record_id":  recordID,
			"message":    "Interview started successfully",
			"start_time": startTime,
		})
		if err != nil {
			log.Printf("Failed to send initial SSE event: sessionID=%s, error=%v", session.SessionID, err)
			return
		}
		// 创建面试引擎，负责执行面试流程，生成问题和反馈，并通过管道写入 SSE 流
		engine := mianshi.NewInterviewEngine(manager, interviewService, writer)
		// 运行面试循环，持续生成问题和反馈，直到面试结束或被取消
		engine.RunInterviewLoop(cancel, session)
	}()
}

// SubmitMianshiAnswer .
// @router /api/mianshi/answer/submit [POST]
func SubmitMianshiAnswer(ctx context.Context, c *app.RequestContext) {
	var err error
	var req mianshiapi.MianshiSubmitInterviewAnswerRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	if req.SessionID == "" {
		response.BadRequest(ctx, c, "session_id is required")
		return
	}
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(ctx, c, "Unauthorized")
		return
	}
	manager := mianshi.GetGlobalSessionManager()
	session := manager.GetSession(req.SessionID)
	if session.UserID != userID {
		response.Unauthorized(ctx, c, "Unauthorized")
		return
	}
	answerToSubmit := req.Answer
	if req.Action != nil && (*req.Action == "quit" || *req.Action == "continue") {
		answerToSubmit = *req.Action
	}
	if err := manager.SubmitAnswer(req.SessionID, answerToSubmit); err != nil {
		response.InternalServerError(ctx, c, err.Error())
	}

	questionIndex := session.QuestionCount
	isLastQuesion := false

	resp := &mianshiapi.MianshiSubmitInterviewAnswerResponse{
		Status:         "received",
		Message:        &req.Answer,
		SessionID:      &req.SessionID,
		QuestionIndex:  &questionIndex,
		IsLastQuestion: &isLastQuesion,
	}

	response.Success(ctx, c, *resp)
}

// GetSession .
// @router /api/mianshi/session/info [GET]
func GetSession(ctx context.Context, c *app.RequestContext) {
	var err error
	var req mianshiapi.MianshiGetSessionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(ctx, c, "Unauthorized")
		return
	}
	sm := mianshi.GetGlobalSessionManager()
	session := sm.GetSession(req.SessionID)
	if session == nil {
		response.NotFound(ctx, c, "Session Not Found")
		return
	}
	if session.UserID != userID {
		response.Unauthorized(ctx, c, "Unauthorized")
		return
	}

	// 构建完的会话信息
	interviewSession := &mianshiapi.InterviewSession{
		SessionID:  session.SessionID,
		UserID:     int32(session.UserID),
		RecordID:   int64(session.RecordID),
		Type:       session.Type,
		Domain:     session.Domain,
		Difficulty: session.Difficulty,
		ResumeID:   &session.ResumeID,
		HasResume:  &session.HasResume,
		StartTime:  session.StartTime.UnixMilli(),
		Status:     session.Status,
	}

	if session.StartTime != (time.Time{}) {
		endTime := session.LastActivity.UnixMilli()
		interviewSession.EndTime = &endTime
	}

	duration := int64(time.Since(session.StartTime).Seconds())
	currentQuestionIndex := session.QuestionCount
	answeredCount := session.QuestionCount
	totalCount := session.QuestionCount

	resp := &mianshiapi.MianshiGetSessionResponse{
		Session:              interviewSession,
		CurrentQuestionIndex: &currentQuestionIndex,
		AnsweredCount:        &answeredCount,
		TotalCount:           &totalCount,
		ElapsedTime:          &duration,
	}

	response.Success(ctx, c, resp)
}

// EndMianshi .
// @router /api/mianshi/interview/end [POST]
func EndMianshi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req mianshiapi.MianshiEndInterviewRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(ctx, c, "Unauthorized")
		return
	}
	sm := mianshi.GetGlobalSessionManager()
	session := sm.GetSession(req.SessionID)
	if session == nil {
		response.NotFound(ctx, c, "Session Not Found")
		return
	}
	if session.UserID != userID {
		response.Unauthorized(ctx, c, "Unauthorized")
		return
	}

	session.Status = "completed"
	endTime := time.Now()
	session.LastActivity = endTime

	duration := int64(time.Since(session.StartTime).Seconds())
	endTimeMs := endTime.UnixMilli()
	totalQuestions := session.QuestionCount
	answeredQuestions := session.QuestionCount

	// 取消面试循环,立即关闭 SSE 连接
	if session.CancelFunc != nil {
		session.CancelFunc()
	}

	interviewService := interviewservice.NewInterviewService()
	updateDTO := &interviews.InterviewRecordDTO{
		ID:       int64(session.RecordID),
		UserID:   int32(session.UserID),
		Status:   "completed",
		Duration: &duration,
	}

	if err = interviewService.UpdateInterviewRecord(ctx, updateDTO); err != nil {
		response.InternalServerError(ctx, c, "Failed to interview record:"+err.Error())
		return
	}

	go func() {
		time.Sleep(1 * time.Second)
		sm.DeleteSession(req.SessionID)
	}()
	log.Printf("[EndMianshi] Interview ended, sessionID:%s, UserID: %d, duration:%d seconds", req.SessionID, userID, duration)

	msg := "Interview ended successfully"
	resp := &mianshiapi.MianshiEndInterviewResponse{
		Status:            "success",
		Message:           &msg,
		Duration:          &duration,
		EndTime:           &endTimeMs,
		TotalQuestions:    &totalQuestions,
		AnsweredQuestions: &answeredQuestions,
	}

	response.Success(ctx, c, resp)
}

// GetMianshiEvaluation .
// @router /api/mianshi/evaluation [GET]
func GetMianshiEvaluation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req mianshiapi.GetMianshiEvaluationRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(ctx, c, "unauthorized")
		return
	}

	reportID := uint64(req.ReportID)
	interviewservice := interviewservice.NewInterviewService()
	existingEvaluation, err := interviewservice.GetInterviewEvaluation(ctx, userID, reportID)
	if err == nil && existingEvaluation != nil {
		response.Success(ctx, c, existingEvaluation)
		return
	}

	resp, err := evaluation.GenerateRecordEvaluation(ctx, userID, reportID)
	if err != nil {
		response.ErrorFromErr(ctx, c, err)
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, c, "Evaluation response is nil")
		return
	}

	response.Success(ctx, c, resp)
}

// GetMianshiAnswerRecord .
// @router /api/mianshi/answer-record [GET]
func GetMianshiAnswerRecord(ctx context.Context, c *app.RequestContext) {
	var err error
	var req mianshiapi.GetMianshiAnswerRecordRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(mianshiapi.GetMianshiAnswerRecordResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetMianshiRecords .
// @router /api/mianshi/records [GET]
func GetMianshiRecords(ctx context.Context, c *app.RequestContext) {
	var err error
	var req mianshiapi.GetMianshiRecordsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(mianshiapi.GetMianshiRecordsResponse)

	c.JSON(consts.StatusOK, resp)
}
