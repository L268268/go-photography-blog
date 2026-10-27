package services

import (
	commentModel "blog/internal/modules/comment/models"
	CommentRepository "blog/internal/modules/comment/repositories"
	"blog/internal/modules/comment/requests/comments"
	UserRepository "blog/internal/modules/user/repositories"
	"blog/internal/modules/user/responses"
	"errors"
	"time"
)

// CommentService 评论服务
type CommentService struct {
	commentRepository CommentRepository.CommentRepositoryInterface
	userRepository    UserRepository.UserRepositoryInterface
}

// CommentResponse 评论响应
type CommentResponse struct {
	ID        uint           `json:"id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	User      responses.User `json:"user"`
}

// New 创建评论服务
func New() *CommentService {
	return &CommentService{
		commentRepository: CommentRepository.New(),
		userRepository:    UserRepository.New(),
	}
}

// Create 创建评论
func (commentService *CommentService) Create(request comments.StoreRequest, user responses.User) error {
	comment := commentModel.Comment{
		Content:   request.Content,
		UserID:    user.ID,
		ArticleID: request.ArticleID,
	}

	result := commentService.commentRepository.Create(comment)
	if result.ID == 0 {
		return errors.New("评论创建失败")
	}

	return nil
}

// 获取文章评论
func (commentService *CommentService) GetComments(articleID uint) []CommentResponse {
	comments := commentService.commentRepository.FindByArticleID(articleID)
	var response []CommentResponse

	for _, comment := range comments {
		// 获取评论用户信息
		user := commentService.userRepository.FindByID(int(comment.UserID))
		userResponse := responses.ToUser(user)

		response = append(response, CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			User:      userResponse,
		})
	}

	return response
}
