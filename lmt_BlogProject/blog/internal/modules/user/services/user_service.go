package services

import (
	"blog/internal/modules/user/information"
	userModel "blog/internal/modules/user/models"
	UserRepository "blog/internal/modules/user/repositories"
	"blog/internal/modules/user/requests/auth"
	UserResponse "blog/internal/modules/user/responses"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// 用户服务
type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

// 获取用户信息
func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

// 创建用户
func (userService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user userModel.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error hashing the password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)

	newUser := userService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("error on creating the user")
	}

	return UserResponse.ToUser(newUser), nil
}

// 检查用户是否存在
func (userService *UserService) CheckUserExists(email string) bool {
	user := userService.userRepository.FindByEmail(email)
	return user.ID != 0
}

// 用户登录
func (userService *UserService) HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error) {
	var response UserResponse.User
	existsUser := userService.userRepository.FindByEmail(request.Email)

	if existsUser.ID == 0 {
		return response, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existsUser.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New("invalid credentials")
	}

	return UserResponse.ToUser(existsUser), nil
}

// 更新个人信息
func (userService *UserService) UpdateInformation(userID uint, request information.UpdateRequest) (UserResponse.User, error) {
	var response UserResponse.User

	// 查找要更新的用户
	var user userModel.User = userService.userRepository.FindByID(int(userID))

	if user.ID == 0 {
		return response, errors.New("用户不存在")
	}

	// 检查是否需要更新密码
	if request.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
		if err != nil {
			return response, errors.New("密码加密失败")
		}
		user.Password = string(hashedPassword)
	}

	// 更新用户名和密码
	user.Name = request.Name
	user.Email = request.Email

	// 调用存储库的更新方法
	updatedUser := userService.userRepository.Update(user)

	if updatedUser.ID == 0 {
		return response, errors.New("更新用户信息失败")
	}

	return UserResponse.ToUser(updatedUser), nil
}
