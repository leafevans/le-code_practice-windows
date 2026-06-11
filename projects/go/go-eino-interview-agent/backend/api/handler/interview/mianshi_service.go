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
func StartMianshiStream(ctx context.Context, c *app.RequestContext) {
	//
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
		response.Unauthorized(ctx, c, "Unauthorized")
		return
	}

	// 调用服务层开始面试
	interviewService := interviewservice.NewInterviewService()
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
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	// 如果请求中包含简历ID，关联简历和面试记录
	var hasResume bool
	var resumeID int64
	if req.ResumeID != nil && *req.ResumeID > 0 {
		hasResume = true
		resumeID = *req.ResumeID
	}

	// 创建面试会话
	manager := mianshi.GetGlobalSessionManager()
	var companyName, positionName string
	if req.CompanyName != nil {
		companyName = *req.CompanyName
	}
	if req.PositionName != nil {
		positionName = *req.PositionName
	}
	session := manager.CreateSessionWithDetails(userID, recordID, resumeID, hasResume, "", req.Type, req.Domain, req.Difficulty, companyName, positionName)

	// 启动面试流程
	cancel, cancelFunc := context.WithCancel(ctx)
	session.CancelFunc = cancelFunc

	// 启动面试流程的goroutine
	mianshi.SetupSSEResonse(c)

	// 这里可以启动一个新的goroutine来处理面试流程
	pipe, writer := io.Pipe()
	c.SetBodyStream(pipe, -1)

	go func() {
		defer func(pipe *io.PipeWriter) {
			err = writer.Close()
			if err != nil {

			}
		}(writer)
		startTime := session.StartTime.UnixMilli()
		err := mianshi.SendSSEEvent(writer, map[string]any{
			"type":       "session_id",
			"session_id": session.SessionID,
			"record_id":  recordID,
			"message":    "Interview started successfully",
			"start_time": startTime,
		})
		if err != nil {
			return
		}

		engine := mianshi.NewInterviewEngine(manager, interviewService, writer)
		engine.RunInterviewLoop(ctx, session)
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
