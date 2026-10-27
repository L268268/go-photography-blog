package repositories

import (
	commentModel "blog/internal/modules/comment/models"
)

type CommentRepositoryInterface interface {
	Create(comment commentModel.Comment) commentModel.Comment
	FindByArticleID(articleID uint) []commentModel.Comment
	FindByID(id uint) commentModel.Comment
}
