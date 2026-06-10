package interview

import (
	"context"
	"io"

	"ai-eino-interview-agent/api/handler/interview/mianshi"
	interviews "ai-eino-interview-agent/api/model/interviews"
	mianshiapi "ai-eino-interview-agent/api/model/mianshi"
	"ai-eino-interview-agent/api/response"
	"ai-eino-interview-agent/internal/middleware"
	interviewservice "ai-eino-interview-agent/internal/service/interviews"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// StartMianshiStream .
// @router /api/mianshi/stream/start [POST]
func StartMianshiStream(ctx context.Context, c *app.RequestContext) {
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
		defer func() {
			
		}()
	}
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

	resp := new(mianshiapi.MianshiSubmitInterviewAnswerResponse)

	c.JSON(consts.StatusOK, resp)
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

	resp := new(mianshiapi.MianshiGetSessionResponse)

	c.JSON(consts.StatusOK, resp)
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

	resp := new(mianshiapi.MianshiEndInterviewResponse)

	c.JSON(consts.StatusOK, resp)
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

	resp := new(mianshiapi.GetMianshiEvaluationResponse)

	c.JSON(consts.StatusOK, resp)
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
