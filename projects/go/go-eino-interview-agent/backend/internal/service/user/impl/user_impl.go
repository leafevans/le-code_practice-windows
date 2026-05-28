package impl

import (
	userapi "ai-eino-interview-agent/api/model/user"
	"ai-eino-interview-agent/internal/middleware"
	"ai-eino-interview-agent/internal/model"
	"context"
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type UserServer struct {
	httpClient *http.Client
}

func NewUserServer() *UserServer {
	return &UserServer{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *UserServer) Register(_ context.Context, req userapi.RegisterRequest) (*userapi.LoginResponse, error) {
	_, err := model.UserDao.FindByUsernameOrEmail(req.GetUsername(), req.GetEmail())
	if err == nil {
		return nil, errors.New("用户名或邮箱已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	userRecord := &model.User{
		Username:     req.GetUsername(),
		Email:        req.GetEmail(),
		PasswordHash: req.GetPassword(),
		Role:         "user",
	}

	if err := model.UserDao.Create(userRecord); err != nil {
		return nil, err
	}

	token, err := middleware.GenerateToken(userRecord.ID, userRecord.Username, userRecord.Role)
	if err != nil {
		return nil, err
	}

	return s.buildLoginResponse(token, userRecord), nil
}

func (s *UserServer) buildLoginResponse(token string, userRecord *model.User) *userapi.LoginResponse {
	resp := userapi.NewLoginResponse()
	resp.Token = token
	resp.User = s.toUserProfile(userRecord)
	return resp
}

func (s *UserServer) toUserProfile(userRecord *model.User) *userapi.UserProfile {
	if userRecord == nil {
		return nil
	}
	profile := userapi.NewUserProfile()
	profile.ID = int64(userRecord.ID)
	profile.Username = userRecord.Username
	profile.Email = userRecord.Email
	profile.Role = userRecord.Role

	if userRecord.WechatOpenID != nil {
		profile.WechatOpenID = userRecord.WechatOpenID
	}
	if userRecord.WechatUnionID != nil {
		profile.WechatUnionID = userRecord.WechatUnionID
	}
	if userRecord.Nickname != "" {
		profile.Nickname = &userRecord.Nickname
	}
	if userRecord.Avatar != "" {
		profile.Avatar = &userRecord.Avatar
	}

	if !userRecord.CreatedAt.IsZero() {
		val := userRecord.CreatedAt.UnixMilli()
		profile.CreatedAt = &val
	}
	if !userRecord.UpdatedAt.IsZero() {
		val := userRecord.UpdatedAt.UnixMilli()
		profile.UpdatedAt = &val
	}
	return profile
}
