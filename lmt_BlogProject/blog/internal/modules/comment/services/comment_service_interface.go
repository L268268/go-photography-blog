package services

import (
	"blog/internal/modules/comment/requests/comments"
	"blog/internal/modules/user/responses"
)

// CommentResponse 评论响应
type CommentServiceInterface interface {
	Create(request comments.StoreRequest, user responses.User) error //创建评论
	GetComments(articleID uint) []CommentResponse                    //获取文章评论
}
