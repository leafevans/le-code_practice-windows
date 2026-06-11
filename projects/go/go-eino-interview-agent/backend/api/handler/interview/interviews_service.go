package interview

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	interviews "ai-eino-interview-agent/api/model/interviews"
	"ai-eino-interview-agent/api/response"
	"ai-eino-interview-agent/chatApp/agent/service"
	"ai-eino-interview-agent/internal/middleware"
	interviewservice "ai-eino-interview-agent/internal/service/interviews"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetInterviewRecords .
// @router /api/interview/records [GET]
func GetInterviewRecords(ctx context.Context, c *app.RequestContext) {
	// 参数绑定和校验
	var err error
	var req interviews.ListInterviewRecordsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		response.BadRequest(ctx, c, err.Error())
		return
	}

	// 用户认证
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(ctx, c, "Unauthorized")
		return
	}

	// 调用服务层获取面试记录
	records, total, err := interviewservice.NewInterviewService().ListInterviewRecords(ctx, userID, req.Page, req.PageSize)
	if err != nil {
		log.Printf("获取面试记录失败: %v", err)
		response.InternalServerError(ctx, c, fmt.Sprintf("获取面试记录失败: %v", err))
		return
	}

	// 构造响应
	resp := new(interviews.ListInterviewRecordsResponse)
	resp.Records = records
	resp.Total = total
	response.Success(ctx, c, resp)
}

// UploadResume .
// @router /api/resume/upload [POST]
func UploadResume(ctx context.Context, c *app.RequestContext) {
	// 用户认证
	userID := middleware.GetUserID(c)

	// 安全校验：确保用户已认证
	if userID == 0 {
		response.Unauthorized(ctx, c, "Authorization token is required")
		return
	}

	// 安全校验：限制上传文件大小和类型
	resumeFilePath, err := handleResumeUpload(c)
	if err != nil {
		response.BadRequest(ctx, c, fmt.Sprintf("上传简历失败: %v", err))
		return
	}
	if resumeFilePath == "" {
		response.BadRequest(ctx, c, "上传简历失败: 无效的文件路径")
		return
	}

	// 文件信息
	fileHeader, err := c.FormFile("resume")
	if err != nil {
		response.BadRequest(ctx, c, fmt.Sprintf("获取文件信息失败: %v", err))
		return
	}

	// 调用简历解析服务
	dbResumeID, _, err := service.ParseResumeAndSave(ctx, userID, resumeFilePath, fileHeader.Size)
	if err != nil {
		log.Printf("解析简历失败: %v", err)
		response.BadRequest(ctx, c, fmt.Sprintf("解析简历失败: %v", err))
		return
	}

	// 构造响应
	resp := &interviews.UploadResumeResponse{
		ResumeID: int64(dbResumeID),
		Message:  "简历上传并解析成功",
	}

	response.Success(ctx, c, resp)
}

// GetResume .
// @router /api/resume/:resume_id [GET]
func GetResume(ctx context.Context, c *app.RequestContext) {
	// 参数绑定和校验
	var req interviews.GetResumeRequest

	// 绑定和校验请求参数
	err := c.BindAndValidate(&req)
	if err != nil {
		response.BadRequest(ctx, c, err.Error())
		return
	}

	// 用户认证
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(ctx, c, "Authorization token is required")
		return
	}

	// 调用服务层获取简历信息
	resumeService := interviewservice.NewResumeService()

	// 安全校验：确保用户只能访问自己的简历
	data, err := resumeService.GetResumeInfoByID(ctx, uint64(req.ResumeID))
	if err != nil {
		log.Printf("获取简历信息失败: %v", err)
		response.InternalServerError(ctx, c, fmt.Sprintf("获取简历信息失败: %v", err))
		return
	}

	// 数据类型转换
	resumeInfo, ok := data.(*interviews.ResumeInfo)
	if !ok {
		log.Printf("数据类型转换失败")
		response.InternalServerError(ctx, c, "数据类型转换失败")
		return
	}
	resp := &interviews.GetResumeResponse{Resume: resumeInfo}

	response.Success(ctx, c, resp)
}

// GetUserResumes .
// @router /api/resume/list [GET]
func GetUserResumes(ctx context.Context, c *app.RequestContext) {
	var req interviews.GetUserResumesRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		response.BadRequest(ctx, c, err.Error())
		return
	}

	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(ctx, c, "Authorization token is required")
		return
	}

	page := int32(1)
	if req.Page != nil && *req.Page > 0 {
		page = *req.Page
	}

	pageSize := int32(10)
	if req.PageSize != nil && *req.PageSize > 0 {
		pageSize = *req.PageSize
	}

	resumeService := interviewservice.NewResumeService()

	data, total, err := resumeService.ListResumeInfosByUserID(ctx, userID, page, pageSize)
	if err != nil {
		log.Printf("获取用户简历列表失败: %v", err)
		response.InternalServerError(ctx, c, fmt.Sprintf("获取用户简历列表失败: %v", err))
		return
	}

	resumeInfo, ok := data.([]*interviews.ResumeInfo)
	if !ok {
		log.Printf("数据类型转换失败")
		response.InternalServerError(ctx, c, "数据类型转换失败")
		return
	}
	resp := &interviews.GetUserResumesResponse{
		Resumes:  resumeInfo,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}

	response.Success(ctx, c, resp)
}

// UpdateResume .
// @router /api/resume/:resume_id [PUT]
func UpdateResume(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interviews.UpdateResumeRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		response.BadRequest(ctx, c, err.Error())
		return
	}

	resp := new(interviews.UpdateResumeResponse)

	c.JSON(consts.StatusOK, resp)
}

// DeleteResume .
// @router /api/resume/:resume_id [DELETE]
func DeleteResume(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interviews.DeleteResumeRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(ctx, c, "Authorization token is required")
		return
	}

	resumeService := interviewservice.NewResumeService()
	err = resumeService.DeleteResume(ctx, userID, uint64(req.ResumeID))
	if err != nil {
		log.Printf("删除简历失败: %v", err)
		response.InternalServerError(ctx, c, fmt.Sprintf("删除简历失败: %v", err))
		return
	}

	resp := &interviews.DeleteResumeResponse{Message: "简历删除成功"}
	response.Success(ctx, c, resp)
}

func handleResumeUpload(requestContext *app.RequestContext) (string, error) {
	// 获取上传的文件夹头
	fileHeader, err := requestContext.FormFile("resume")
	if err != nil {
		return "", err
	}

	if fileHeader.Filename == "" || !strings.HasSuffix(fileHeader.Filename, "pdf") {
		return "", errors.New("无效的文件类型，请上传PDF格式的简历")
	}

	// 安全校验文件大小（示例：限制为10MB）
	if fileHeader.Size > 10*1024*1024 {
		return "", errors.New("文件大小超出限制，请上传小于10MB的PDF文件")
	}

	// 打开上传的文件
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			// 记录关闭文件时的错误，但不影响主流程
			log.Printf("关闭上传文件时发生错误: %v", err)
		}
	}(file)

	// 生成文件目录
	projectRoot := filepath.Join(os.Getenv("PWD"), "uploads", "resumes")

	// 生成唯一文件名
	timestamp := time.Now().UnixNano()
	filename := fmt.Sprintf("resume_%d_%s", timestamp, fileHeader.Filename)
	resumeFilePath := filepath.Join(projectRoot, filename)

	// 创建本地文件
	tempFile, err := os.Create(resumeFilePath)
	if err != nil {
		return "", err
	}

	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {
			log.Printf("关闭临时文件时发生错误: %v", err)
		}
	}(tempFile)

	// 复制文件内容到本地文件
	if _, err := io.Copy(tempFile, file); err != nil {
		err := os.Remove(resumeFilePath) // 删除已创建的文件
		return "", fmt.Errorf("保存文件失败: %w", err)
	}

	// 强制刷盘（确保文件写入磁盘）
	if err := tempFile.Sync(); err != nil {
		return "", fmt.Errorf("刷盘失败: %w", err)
	}

	return resumeFilePath, nil
}
