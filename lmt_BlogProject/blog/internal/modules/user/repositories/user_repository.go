package repositories

import (
	userModel "blog/internal/modules/user/models"
	"blog/pkg/database"

	"gorm.io/gorm"
)

// UserRepository 用户仓库接口
type UserRepository struct {
	DB *gorm.DB
}

// New 创建用户仓库实例
func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

// Create 创建用户
func (userRepository *UserRepository) Create(user userModel.User) userModel.User {
	var newUser userModel.User

	userRepository.DB.Create(&user).Scan(&newUser)

	return newUser
}

// FindByEmail 根据邮箱查找用户
func (userRepository *UserRepository) FindByEmail(email string) userModel.User {
	var user userModel.User

	userRepository.DB.First(&user, "email = ?", email)

	return user
}

// FindAll 查找所有用户
func (userRepository *UserRepository) FindByID(id int) userModel.User {
	var user userModel.User

	userRepository.DB.First(&user, "id = ?", id)

	return user
}

// Update 更新用户信息
func (userRepository *UserRepository) Update(user userModel.User) userModel.User {
	// 更新用户信息
	result := userRepository.DB.Save(&user)
	if result.Error != nil {
		return userModel.User{}
	}

	return user //更新成功，返回用户信息
}
