package tool

import (
	"ai-eino-interview-agent/internal/model"
	"context"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

// GetMianshiInfoRequest 获取面试对话记录的请求结构体
type GetMianshiInfoRequest struct {
	UserID   uint   `json:"user_id" jsonschema:"description=用户ID"`
	ReportID uint64 `json:"report_id" jsonschema:"description=报告ID"`
}

// GetMianshiInfoResponse 获取面试对话记录的响应结构体
type GetMianshiInfoResponse struct {
	Data []model.InterviewDialogue `json:"data" jsonschema:"description=面试对话记录列表"`
}

// GetMianshiInfo 获取面试的完整问题和对话记录
func GetMianshiInfo(ctx context.Context, req *GetMianshiInfoRequest) (*GetMianshiInfoResponse, error) {
	// 如果请求参数为空，直接返回空响应
	if req == nil {
		return nil, nil
	}

	// 从数据库中获取面试对话记录
	dialogues, err := model.InterviewDialogueDao.GetInterviewDialoguesByUserIdAndRecordId(req.UserID, req.ReportID)
	if err != nil {
		return nil, err
	}

	// 将获取到的对话记录封装到响应结构体中返回
	return &GetMianshiInfoResponse{
		Data: *dialogues,
	}, nil
}

// GetMianshiInfoTool 定义获取面试信息的工具
func GetMianshiInfoTool() tool.InvokableTool {
	inferTool, err := utils.InferTool(
		"get_mianshi_info",
		"获取面试的完整问题和对话记录",
		GetMianshiInfo)
	if err != nil {
		panic(err)
	}
	return inferTool
}
