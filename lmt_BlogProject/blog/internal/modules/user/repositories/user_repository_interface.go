package repositories

import (
	userModel "blog/internal/modules/user/models"
)

// 定义用户仓库接口
type UserRepositoryInterface interface {
	Create(user userModel.User) userModel.User //创建用户的接口
	FindByEmail(email string) userModel.User   //根据email查找用户
	FindByID(id int) userModel.User            //根据id查找用户
	Update(user userModel.User) userModel.User //更新用户的接口
}
