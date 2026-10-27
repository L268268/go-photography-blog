package articles

// UpdateRequest 更新文章请求
type UpdateRequest struct {
	Title   string `form:"title" binding:"required,min=3,max=200"`
	Content string `form:"content" binding:"required,min=10"`
}
