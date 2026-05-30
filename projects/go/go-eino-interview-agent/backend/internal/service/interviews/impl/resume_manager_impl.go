package impl

import (
	"ai-eino-interview-agent/api/model/interviews"
	"ai-eino-interview-agent/internal/model"
	"context"
	"fmt"
	"log"
)

type ResumeServer struct{}

func NewResumeServer() *ResumeServer {
	return &ResumeServer{}
}

func toResumeInfo(dataMap map[string]any) *interviews.ResumeInfo {
	resumeInfo := &interviews.ResumeInfo{}

	// 提取 ID
	if id, ok := dataMap["id"]; ok {
		if idVal, ok := id.(uint64); ok {
			resumeInfo.ID = int64(idVal)
		} else if idVal, ok := id.(int64); ok {
			resumeInfo.ID = idVal
		}
	}

	// 提取 UserID
	if userID, ok := dataMap["user_id"]; ok {
		if userIDVal, ok := userID.(uint); ok {
			resumeInfo.UserID = int32(userIDVal)
		} else if userIDVal, ok := userID.(int); ok {
			resumeInfo.UserID = int32(userIDVal)
		}
	}

	// 提取 FileName
	if fileName, ok := dataMap["file_name"].(string); ok {
		resumeInfo.FileName = fileName
	}

	// 提取 FileSize
	if fileSize, ok := dataMap["file_size"].(int64); ok {
		resumeInfo.FileSize = fileSize
	}

	// 提取 FileType
	if fileType, ok := dataMap["file_type"].(string); ok {
		resumeInfo.FileType = fileType
	}

	// 提取 IsDefault
	if isDefault, ok := dataMap["is_default"]; ok {
		if isDefaultVal, ok := isDefault.(int); ok {
			resumeInfo.IsDefault = int32(isDefaultVal)
		} else if isDefaultVal, ok := isDefault.(int32); ok {
			resumeInfo.IsDefault = isDefaultVal
		}
	}

	// 提取 CreatedAt
	if createdAt, ok := dataMap["created_at"]; ok {
		if createdAtVal, ok := createdAt.(int64); ok {
			resumeInfo.CreatedAt = createdAtVal
		}
	}

	// 提取 UpdatedAt
	if updatedAt, ok := dataMap["updated_at"]; ok {
		if updatedAtVal, ok := updatedAt.(int64); ok {
			resumeInfo.UpdatedAt = updatedAtVal
		}
	}

	return resumeInfo
}

// 上传简历
func (*ResumeServer) UploadResume(_ context.Context, userID uint, fileName string, fileSize int64, fileType string, content string) (uint64, error) {
	resume := &model.Resume{
		UserID:    userID,
		Content:   content,
		FileName:  fileName,
		FileSize:  fileSize,
		FileType:  fileType,
		IsDefault: 1,
		Deleted:   0,
	}
	resumeID, err := model.ResumeDao.Create(resume)
	if err != nil {
		log.Printf("创建简历记录失败: %v", err)
		return 0, err
	}
	log.Printf("简历上传成功，记录ID: %d", resumeID)
	return resumeID, nil
}

func (*ResumeServer) DeleteResume(_ context.Context, userID uint, resumeID uint64) error {
	resume, err := model.ResumeDao.GetResumeByID(resumeID)
	if err != nil {
		return err
	}
	if resume.UserID != userID {
		log.Printf("用户 %d 无权删除简历 %d", userID, resumeID)
		return fmt.Errorf("无权删除该简历")
	}
	err = model.ResumeDao.DeleteResume(resumeID)
	if err != nil {
		log.Printf("删除简历记录失败: %v", err)
		return err
	}
	log.Printf("简历删除成功，记录ID: %d", resumeID)
	return nil
}

func (*ResumeServer) GetResumeInfoByID(_ context.Context, userID uint, resumeID uint64) (any, error) {
	resume, err := model.ResumeDao.GetResumeByID(resumeID)
	if err != nil {
		return nil, err
	}
	dataMap := map[string]any{
		"id":         resume.ID,
		"user_id":    resume.UserID,
		"file_name":  resume.FileName,
		"file_size":  resume.FileSize,
		"file_type":  resume.FileType,
		"is_default": resume.IsDefault,
		"created_at": resume.CreatedAt,
		"updated_at": resume.UpdatedAt,
	}

	return toResumeInfo(dataMap), nil
}

// 获取用户简历列表分页
func (*ResumeServer) ListResumeInfosByUserID(
	_ context.Context,
	userID uint,
	page, pageSize int32,
) (any, int64, error) {
	resumes, total, err := model.ResumeDao.ListByUser(userID, page, pageSize)
	if err != nil {
		log.Printf("获取用户简历列表失败: %v", err)
		return nil, 0, err
	}
	var resumeInfos []*interviews.ResumeInfo
	for _, resume := range resumes {
		dataMap := map[string]any{
			"id":         resume.ID,
			"user_id":    resume.UserID,
			"file_name":  resume.FileName,
			"file_size":  resume.FileSize,
			"file_type":  resume.FileType,
			"is_default": resume.IsDefault,
			"created_at": resume.CreatedAt,
			"updated_at": resume.UpdatedAt,
		}
		resumeInfos = append(resumeInfos, toResumeInfo(dataMap))
	}
	return resumeInfos, total, nil
}
