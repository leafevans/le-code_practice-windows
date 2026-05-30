package service

import (
	"ai-eino-interview-agent/chatApp/agent/resume"
	"ai-eino-interview-agent/internal/model"
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

// ResumeParseResult 简历解析结果
type ResumeParseResult struct {
	BasicInfo struct {
		Name      string `json:"name"`
		WorkYears string `json:"work_years"`
		Contact   string `json:"contact"`
	} `json:"basic_info"`
	Education []struct {
		School         string `json:"school"`
		Major          string `json:"major"`
		Degree         string `json:"degree"`
		GraduationYear string `json:"graduation_year"`
	} `json:"education"`
	WorkExperience []struct {
		Company          string `json:"company"`
		Position         string `json:"position"`
		Duration         string `json:"duration"`
		Responsibilities string `json:"responsibilities"`
	} `json:"work_experience"`
	TechStack                   []string      `json:"tech_stack"`
	Projects                    []interface{} `json:"projects"`
	Skills                      []string      `json:"skills"`
	Certifications              []string      `json:"certifications"`
	Strengths                   string        `json:"strengths"`
	PotentialWeaknesses         string        `json:"potential_weaknesses"`
	RecommendedDifficulty       string        `json:"recommended_difficulty"`
	InterviewFocusAreas         []string      `json:"interview_focus_areas"`
	SuggestedQuestionDirections []string      `json:"suggested_questions_directions"`
}

// ParseResumeAndSave 解析简历并保存到数据库
func ParseResumeAndSave(ctx context.Context, userID uint, resumeFilePath string, filesize int64) (uint64, *ResumeParseResult, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	agent, err := resume.NewResumeAgent(userID)
	if err != nil {
		return 0, nil, fmt.Errorf("创建简历解析智能体失败：%w", err)
	}

	runner := adk.NewRunner(timeoutCtx, adk.RunnerConfig{
		Agent: agent,
	})

	query := fmt.Sprintf(`【重要】请立即解析以下简历文件并提取关键信息：

简历文件路径：%s

【必须执行的步骤】：
1. 【第一步】立即使用 pdf_to_text 工具解析简历文件，获取完整的简历文本内容
2. 【第二步】从解析的简历文本中提取所有关键信息（姓名、工作年限、联系方式、教育背景、工作经历、技术栈、项目经验、技能、证书等）
3. 【第三步】分析候选人的背景特点和核心竞争力
4. 【第四步】生成面试建议和推荐难度

【重要提示】：
- 不要跳过 pdf_to_text 工具调用
- 必须从简历内容中提取真实的信息，不要返回空数据
- 所有JSON字段都必须填充实际内容
- 只返回JSON格式，不要返回其他文本

请返回完整的 JSON 格式结果。`, resumeFilePath)

	userMsg := &schema.Message{
		Role:    schema.User,
		Content: query,
	}

	messages := []adk.Message{userMsg}
	iter := runner.Run(timeoutCtx, messages)

	var lastMessage string

	for {
		if timeoutCtx.Err() != nil {
			log.Fatal("简历解析超时")
			return 0, nil, fmt.Errorf("简历解析超时")
		}

		event, ok := iter.Next()
		if !ok {
			break
		}

		if event.Err != nil {
			log.Printf("简历解析过程中发生错误: %v", event.Err)
			return 0, nil, fmt.Errorf("简历解析过程中发生错误: %w", event.Err)
		}

		if event.Output != nil && event.Output.MessageOutput != nil {
			lastMessage = event.Output.MessageOutput.Message.Content
		}
	}

	if lastMessage == "" {
		log.Println("简历解析失败，未收到智能体响应")
		return 0, nil, fmt.Errorf("简历解析失败，未收到智能体响应")
	}

	log.Printf("智能体响应: %s", lastMessage)
	parseResult := parseResumeResponse(lastMessage)
	if parseResult == nil {
		log.Println("简历解析失败，无法解析智能体响应")
		return 0, nil, fmt.Errorf("简历解析失败，无法解析智能体响应")
	}

	if !isValidResumeResult(parseResult) {
		log.Println("简历解析结果无效，缺少必要字段")
		return 0, nil, fmt.Errorf("简历解析结果无效，缺少必要字段")
	}

	resumeID, err := saveResumeResult(ctx, userID, parseResult, filesize, resumeFilePath)
	if err != nil {
		log.Printf("保存简历结果失败: %v", err)
		return 0, nil, fmt.Errorf("保存简历结果失败: %w", err)
	}

	return resumeID, parseResult, nil
}

func saveResumeResult(_ context.Context, userID uint, result *ResumeParseResult, filesize int64, resumeFilePath string) (uint64, error) {
	content, err := sonic.Marshal(result)
	if err != nil {
		return 0, fmt.Errorf("序列化简历解析结果失败: %w", err)
	}

	fileName := filepath.Base(resumeFilePath)
	log.Printf("保存简历结果，文件名: %s, 文件大小: %d", fileName, filesize)

	resumeRecord := &model.Resume{
		UserID:    userID,
		Content:   string(content),
		FileName:  fileName,
		FileSize:  filesize,
		FileType:  "pdf",
		IsDefault: 1,
		Deleted:   0,
	}

	resumeID, err := model.ResumeDao.Create(resumeRecord)
	if err != nil {
		return 0, fmt.Errorf("创建简历记录失败: %w", err)
	}
	return resumeID, nil
}

func isValidResumeResult(result *ResumeParseResult) bool {
	if result == nil {
		return false
	}

	// 检查基本信息是否有内容
	if result.BasicInfo.Name != "" || result.BasicInfo.WorkYears != "" || result.BasicInfo.Contact != "" {
		return true
	}

	// 检查教育背景
	if len(result.Education) > 0 {
		return true
	}

	// 检查工作经历
	if len(result.WorkExperience) > 0 {
		return true
	}

	// 检查技术栈
	if len(result.TechStack) > 0 {
		return true
	}

	// 检查项目经验
	if len(result.Projects) > 0 {
		return true
	}

	// 检查技能
	if len(result.Skills) > 0 {
		return true
	}

	// 检查证书
	if len(result.Certifications) > 0 {
		return true
	}

	// 检查其他字段
	if result.Strengths != "" || result.PotentialWeaknesses != "" {
		return true
	}

	// 检查面试关注领域
	if len(result.InterviewFocusAreas) > 0 {
		return true
	}

	// 检查建议的提问方向
	if len(result.SuggestedQuestionDirections) > 0 {
		return true
	}

	// 如果所有字段都是空的，返回 false
	return false
}

func parseResumeResponse(text string) *ResumeParseResult {
	result := &ResumeParseResult{}
	if err := sonic.Unmarshal([]byte(text), result); err != nil {
		log.Printf("直接解析简历响应失败: %v", err)

		jsonStr := ExtractJSONFromMessage(text)
		if jsonStr == "" {
			log.Println("从智能体响应中提取JSON失败")
			return nil
		}

		log.Printf("提取的JSON字符串: %s", jsonStr)
		err := sonic.Unmarshal([]byte(jsonStr), result)
		if err != nil {
			log.Printf("解析提取的JSON失败: %v", err)
			return nil
		}
	}

	return result
}

// ExtractJSONFromMessage 从智能体响应中提取纯 JSON 内容
func ExtractJSONFromMessage(text string) string {
	start := -1
	for i := 0; i < len(text); i++ {
		if text[i] == '{' || text[i] == '[' {
			start = i
			break
		}
	}
	if start == -1 {
		return ""
	}

	var stack []byte
	inString := false
	escaped := false

	for i := start; i < len(text); i++ {
		ch := text[i]

		if inString {
			if escaped {
				escaped = false
				continue
			}
			if ch == '\\' {
				escaped = true
				continue
			}
			if ch == '"' {
				inString = false
			}
			continue
		}

		switch ch {
		case '"':
			inString = true
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != ch {
				return ""
			}
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				return text[start : i+1]
			}
		}
	}

	return ""
}
