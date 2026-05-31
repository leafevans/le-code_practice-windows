package impl

import (
	interviewsapi "ai-eino-interview-agent/api/model/interviews"
	"ai-eino-interview-agent/internal/model"
	"context"
	"fmt"
	"log"
)

// InterviewServiceImpl 是面试服务的实现结构体
type InterviewServiceImpl struct{}

// NewInterviewService 创建一个新的面试服务实例
func NewInterviewServiceImpl() *InterviewServiceImpl {
	return &InterviewServiceImpl{}
}

// CreateInterviewRecord 创建面试记录
func (s *InterviewServiceImpl) CreateInterviewRecord(ctx context.Context, dto *interviewsapi.InterviewRecordDTO) (uint64, error) {
	// 处理 DTO 中的可选字段，确保不会出现 nil 指针异常
	companyName := ""
	if dto.CompanyName != nil {
		companyName = *dto.CompanyName
	}

	// 处理职位名称字段，确保不会出现 nil 指针异常
	positionName := ""
	if dto.PositionName != nil {
		positionName = *dto.PositionName
	}

	// 处理状态字段，设置默认值为 "pending" 如果 DTO 中没有提供状态
	status := dto.Status
	if status == "" {
		status = "pending"
	}

	// 	处理持续时间字段，确保不会出现 nil 指针异常
	var duration int64 = 0
	if dto.Duration != nil {
		duration = *dto.Duration
	}

	// 构建 InterviewRecord 实例
	record := &model.InterviewRecord{
		UserID:       uint(dto.UserID),
		Type:         dto.Type,
		Difficulty:   dto.Difficulty,
		Domain:       dto.Domain,
		CompanyName:  companyName,
		PositionName: positionName,
		Status:       status,
		Duration:     duration,
	}

	// 将面试记录保存到数据库，并获取生成的记录ID
	recordID, err := model.InterviewRecordDao.CreateInterviewRecord(record)
	if err != nil {
		log.Printf("Failed to create interview record: %v", err)
		return 0, err
	}

	// 记录日志，包含面试记录的关键信息
	log.Printf("Successfully created interview record: recordID=%d, userID=%d, type=%s, difficulty=%s, domain=%s", recordID, dto.UserID, dto.Type, dto.Difficulty, dto.Domain)
	return recordID, nil
}

// UpdateInterviewRecord 更新面试记录
func (s *InterviewServiceImpl) UpdateInterviewRecord(ctx context.Context, dto *interviewsapi.InterviewRecordDTO) error {
	// 处理 DTO 指针字段，确保不会出现 nil 指针异常
	companyName := ""
	if dto.CompanyName != nil {
		companyName = *dto.CompanyName
	}
	positionName := ""
	if dto.PositionName != nil {
		positionName = *dto.PositionName
	}
	var duration int64 = 0
	if dto.Duration != nil {
		duration = *dto.Duration
	}

	// 构建 InterviewRecord 实例
	record := &model.InterviewRecord{
		ID:           uint64(dto.ID),
		UserID:       uint(dto.UserID),
		Type:         dto.Type,
		Difficulty:   dto.Difficulty,
		Domain:       dto.Domain,
		CompanyName:  companyName,
		PositionName: positionName,
		Status:       dto.Status,
		Duration:     duration,
	}

	log.Printf("Updating interview record: recordID=%d, userID=%d, type=%s, difficulty=%s, domain=%s", record.ID, record.UserID, record.Type, record.Difficulty, record.Domain)

	return model.InterviewRecordDao.UpdateInterviewRecord(record)
}

// ListInterviewRecords 列出面试记录
func (s *InterviewServiceImpl) ListInterviewRecords(ctx context.Context, userID uint, page, pageSize *int32) ([]*interviewsapi.InterviewRecordDTO, int64, error) {
	// 默认页码为1
	pageNum := int32(1)
	if page != nil {
		pageNum = *page
	}

	// 默认每页10条记录
	pageSz := int32(10)
	if pageSize != nil {
		pageSz = *pageSize
	}

	// 从数据库中获取面试记录列表和总记录数
	records, total, err := model.InterviewRecordDao.ListInterviewRecords(userID, &pageNum, &pageSz)
	if err != nil {
		log.Printf("Failed to list interview records for userID=%d: %v", userID, err)
		return nil, 0, err
	}

	// 将 InterviewRecord 转换为 InterviewRecordDTO
	dtoList := make([]*interviewsapi.InterviewRecordDTO, 0, len(records))
	for _, record := range records {
		dto := convertToInterviewRecordDTO(record)
		dtoList = append(dtoList, dto)
	}

	// 记录日志，包含分页信息
	log.Printf("Listed interview records for userID=%d: total=%d, page=%d, pageSize=%d", userID, total, pageNum, pageSz)
	return dtoList, total, nil
}

// convertToInterviewRecordDTO 将 InterviewRecord 转换为 InterviewRecordDTO
func convertToInterviewRecordDTO(record *model.InterviewRecord) *interviewsapi.InterviewRecordDTO {
	dto := interviewsapi.NewInterviewRecordDTO()
	dto.ID = int64(record.ID)
	dto.UserID = int32(record.UserID)
	dto.Type = record.Type
	dto.Difficulty = record.Difficulty
	dto.Domain = record.Domain
	dto.Status = record.Status

	// 处理可选字段，只有在非空时才设置 DTO 字段
	if record.CompanyName != "" {
		v := record.CompanyName
		dto.CompanyName = &v
	}
	if record.PositionName != "" {
		v := record.PositionName
		dto.PositionName = &v
	}
	if record.Duration != 0 {
		v := record.Duration
		dto.Duration = &v
	}

	// 将 CreatedAt 和 UpdatedAt 转换为毫秒时间戳
	if !record.CreatedAt.IsZero() {
		ms := record.CreatedAt.UnixNano() / 1e6
		dto.CreatedAt = &ms
	}

	if !record.UpdatedAt.IsZero() {
		ms := record.UpdatedAt.UnixNano() / 1e6
		dto.UpdatedAt = &ms
	}

	return dto
}

// GetInterviewEvaluation 获取面试评估信息
func (s *InterviewServiceImpl) GetInterviewEvaluation(ctx context.Context, userID uint, reportID uint64) (any, error) {
	// 根据用户ID和面试记录ID获取评估信息
	evaluation, err := model.InterviewEvaluationDao.GetEvaluationByUserIDAndReportID(userID, reportID)
	if err != nil {
		log.Printf("Failed to get interview evaluation for userID=%d, reportID=%d: %v", userID, reportID, err)
		return nil, err
	}

	// 将 InterviewEvaluation 转换为 map[string]any 以适配 API 返回格式
	return map[string]any{
		"id":         evaluation.ID,
		"user_id":    evaluation.UserID,
		"report_id":  evaluation.ReportID,
		"comment":    evaluation.Comment,
		"score":      evaluation.Score,
		"dimensions": evaluation.Dimensions,
		"created_at": evaluation.CreatedAt,
		"updated_at": evaluation.UpdatedAt,
	}, nil
}

// GetAnswerReport 获取面试答案报告
func (s *InterviewServiceImpl) GetAnswerReport(ctx context.Context, userID uint, reportID uint64) (any, error) {
	// 根据用户ID和面试记录ID获取答案报告
	report, err := model.AnswerReportDao.GetAnswerReportByUserIDAndReportID(userID, reportID)
	if err != nil {
		log.Printf("[GetAnswerReport] 获取答题报告失败: %v", err)
		return nil, err
	}

	// 返回答题报告数据
	return map[string]any{
		"id":         report.ID,
		"user_id":    report.UserID,
		"report_id":  report.ReportID,
		"records":    report.Records,
		"deleted":    report.Deleted,
		"created_at": report.CreatedAt,
		"updated_at": report.UpdatedAt,
	}, nil
}

// SaveInterviewDialogueWithParent 保存面试对话（支持父子关系）
func (s *InterviewServiceImpl) SaveInterviewDialogueWithParent(
	ctx context.Context,
	userID uint, reportID uint64,
	mainQuestion *model.InterviewDialogue,
	followUpQuestions []*model.InterviewDialogue) error {
	// 首先保存主问题
	mainQuestion.UserID = userID
	mainQuestion.ReportID = reportID
	if err := model.InterviewDialogueDao.Create(mainQuestion); err != nil {
		return fmt.Errorf("failed to save main question: %w", err)
	}

	// 然后保存所有的追问，确保它们都关联到同一个用户ID和报告ID
	for _, followUpQuestion := range followUpQuestions {
		followUpQuestion.UserID = userID
		followUpQuestion.ReportID = reportID
		if err := model.InterviewDialogueDao.Create(followUpQuestion); err != nil {
			return fmt.Errorf("failed to save follow-up question: %w", err)
		}
	}

	return nil
}
