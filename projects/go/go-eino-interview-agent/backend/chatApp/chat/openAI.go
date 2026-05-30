package chat

import (
	"context"
	"strings"

	"ai-eino-interview-agent/internal/errors"
	usermodel "ai-eino-interview-agent/internal/model"
	"ai-eino-interview-agent/internal/service/common"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
)

func CreateOpenAIChatModel(ctx context.Context, userID uint) (model.ToolCallingChatModel, error) {
	result, err := usermodel.UserModelDao.GetDefaultUserModel(int64(userID))
	if err != nil {
		return nil, errors.NewDBError("Failed to get user model", err)
	}

	apiKey, err := common.DecryptAPIKey(result.APIKeyEncrypted)
	if err != nil {
		return nil, errors.NewInternalError("Failed to decrypt API key", err)
	}

	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  apiKey,
		Model:   result.ModelKey,
		BaseURL: result.BaseURL,
	})
	if err != nil {
		errMsg := strings.ToLower(err.Error())
		switch {
		case strings.Contains(errMsg, "insufficient_quota") ||
			strings.Contains(errMsg, "billing_not_active") ||
			strings.Contains(errMsg, "quota_exceeded") ||
			strings.Contains(errMsg, "insufficient tokens"):
			return nil, errors.NewInsufficientTokensError("Model API: Insufficient tokens or quota exceeded. Please check your account balance.", err)
		case strings.Contains(errMsg, "rate_limit_exceeded") ||
			strings.Contains(errMsg, "too_many_requests") ||
			strings.Contains(errMsg, "rate limit"):
			return nil, errors.NewRateLimitExceededError("Model API: Rate limit exceeded. Please try again later.", err)
		case strings.Contains(errMsg, "context_length_exceeded") ||
			strings.Contains(errMsg, "maximum context length") ||
			strings.Contains(errMsg, "token limit"):
			return nil, errors.NewContextLengthExceededError("Model API: Context length exceeded. Please try with shorter input.", err)
		default:
			return nil, errors.NewOpenAIError("Failed to create OpenAI chat model", err)
		}
	}

	return chatModel, nil
}
