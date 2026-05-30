package tool

import (
	"ai-eino-interview-agent/internal/model"
	"context"
	"log"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

// GetResumeInfoRequest 获取简历信息请求体
type GetResumeInfoRequest struct {
	ResumeID uint64 `json:"resume_id" jsonschema:"description=简历ID"`
}

// GetResumeInfoResponse 获取面试数据的响应结构体
type GetResumeInfoResponse struct {
	Data *model.Resume `json:"data" jsonschema:"description=简历解析后的内容"`
}

// 获取简历数据
func GetResumeInfo(ctx context.Context, request *GetResumeInfoRequest) (*GetResumeInfoResponse, error) {
	if request == nil {
		return nil, nil
	}

	resume, err := model.ResumeDao.GetResumeByID(request.ResumeID)
	if err != nil {
		log.Printf("获取简历信息失败: %v", err)
		return nil, err
	}

	return &GetResumeInfoResponse{Data: resume}, nil
}

// 创建获取简历信息工具
func GetResumeInfoTool() tool.InvokableTool {
	inferTool, err := utils.InferTool(
		"get_resume_info",
		"获取简历信息",
		GetResumeInfo,
	)
	if err != nil {
		log.Fatalf("创建获取简历信息工具失败: %v", err)
	}
	return inferTool
}
