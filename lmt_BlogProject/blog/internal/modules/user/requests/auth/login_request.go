package auth

// LoginRequest 登录请求
type LoginRequest struct {
	Email    string `form:"email" binding:"required,email,min=3,max=100"`
	Password string `form:"password" binding:"required,min=8"`
}
