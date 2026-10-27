package repositories

import (
	commentModel "blog/internal/modules/comment/models"
	"blog/pkg/database"

	"gorm.io/gorm"
)

// 评论仓库接口
type CommentRepository struct {
	DB *gorm.DB
}

// 初始化评论仓库
func New() *CommentRepository {
	return &CommentRepository{
		DB: database.Connection(),
	}
}

// 创建新的评论
func (commentRepository *CommentRepository) Create(comment commentModel.Comment) commentModel.Comment {
	result := commentRepository.DB.Create(&comment)
	if result.Error != nil {
		return commentModel.Comment{}
	}
	return comment
}

// 根据文章ID查找评论
func (commentRepository *CommentRepository) FindByArticleID(articleID uint) []commentModel.Comment {
	var comments []commentModel.Comment
	commentRepository.DB.Where("article_id = ?", articleID).Order("created_at desc").Find(&comments)
	return comments
}

func (commentRepository *CommentRepository) FindByID(id uint) commentModel.Comment {
	var comment commentModel.Comment
	commentRepository.DB.First(&comment, id)
	return comment
}
