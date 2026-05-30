package interviews

import (
	"ai-eino-interview-agent/internal/service/interviews/impl"
	"context"
)

// NewResumeService 返回简历服务的实现
func NewResumeService() ResumeManager {
	return impl.NewResumeServer()
}

// ResumeManager 简历管理接口
type ResumeManager interface {
	// UploadResume 上传简历，返回简历ID
	UploadResume(
		ctx context.Context,
		userID uint,
		fileName string,
		fileType string,
		fileSize int64,
		content string,
	) (uint64, error)

	// DeleteResume 删除简历
	DeleteResume(
		ctx context.Context,
		userID uint,
		resumeID uint64,
	) error

	// GetResumeInfoByID 根据简历ID获取简历详情，返回强类型 ResumeInfo
	GetResumeInfoByID(
		ctx context.Context,
		resumeID uint64,
	) (any, error)

	// ListResumeInfosByUserID 分页获取用户的简历列表，返回强类型 ResumeInfo 列表
	ListResumeInfosByUserID(
		ctx context.Context,
		userID uint,
		page, pageSize int32,
	) (any, int64, error)
}
