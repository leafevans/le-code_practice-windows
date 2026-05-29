package user

import (
	userapi "ai-eino-interview-agent/api/model/user"
	"ai-eino-interview-agent/internal/model"
	"ai-eino-interview-agent/internal/service/user/impl"
	"context"
)

// NewModelManager 返回用户模型管理接口的默认实现
func NewModelManager() ModelManager {
	return impl.NewUserModelServer()
}

// NewUserManager 返回用户管理接口的默认实现
func NewUserManager() UserManager {
	return impl.NewUserServer()
}

type ModelManager interface {
	CreateUserModel(
		ctx context.Context,
		userID int64,
		req userapi.CreateUserModelRequest,
	) (string, error)
	ListUserModels(
		ctx context.Context,
		userID int64,
		req userapi.ListUserModelsRequest,
	) ([]*model.UserModel, int64, error)
	UserModelDetail(
		ctx context.Context,
		userID int64,
		modelID int64,
	) (*model.UserModel, error)
	UpdateUserModel(
		ctx context.Context,
		userID int64,
		req userapi.UpdateUserModelRequest,
	) error
	DeleteUserModel(
		ctx context.Context,
		userID int64,
		modelID int64,
	) error
	// CheckUserModelConfigured 检查用户是否配置了默认模型
	CheckUserModelConfigured(
		ctx context.Context,
		userID int64,
	) (*model.UserModel, error)
}

type UserManager interface {
	Register(ctx context.Context, req userapi.RegisterRequest) (*userapi.LoginResponse, error)
	Login(ctx context.Context, req userapi.LoginRequest) (*userapi.LoginResponse, error)
}
