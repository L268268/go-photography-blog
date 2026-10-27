package comments

// StoreRequest 创建评论请求
type StoreRequest struct {
	Content   string `form:"content" binding:"required,min=1"`
	ArticleID uint   `form:"article_id" binding:"required"`
}
