package services

import (
	"blog/internal/modules/user/information"
	"blog/internal/modules/user/requests/auth"
	UserResponse "blog/internal/modules/user/responses"
)

// UserServiceInterface 用户服务接口
type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
	CheckUserExists(email string) bool
	HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
	UpdateInformation(userID uint, request information.UpdateRequest) (UserResponse.User, error) //更新用户信息的函数接口
}
