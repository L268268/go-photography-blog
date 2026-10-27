package information

// UpdateRequest 更新用户信息请求
type UpdateRequest struct {
	Name     string `form:"name" binding:"required,min=1,max=100"`
	Email    string `form:"email" binding:"required,email,min=3,max=100"`
	Password string `form:"password" binding:"required,min=8"`
}
