package repositories

import (
	ArticleModel "blog/internal/modules/article/models"
	"blog/pkg/database"
	"fmt"

	"gorm.io/gorm"
)

// ArticleRepository 文章存储库
type ArticleRepository struct {
	DB *gorm.DB
}

// New 创建新的文章存储库
func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

// List 获取文章列表
func (articleRepository *ArticleRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article

	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)

	return articles
}

// FindByID 根据ID查找文章
func (articleRepository *ArticleRepository) Find(id int) ArticleModel.Article {
	var article ArticleModel.Article

	articleRepository.DB.Joins("User").First(&article, id)

	return article
}

// Create 创建文章
func (articleRepository *ArticleRepository) Create(article ArticleModel.Article) ArticleModel.Article {
	var newArticle ArticleModel.Article

	articleRepository.DB.Create(&article).Scan(&newArticle)

	return newArticle
}

// Update 更新文章
func (repository *ArticleRepository) Update(article ArticleModel.Article) ArticleModel.Article {

	result := database.DB.Save(&article)
	if result.Error != nil {
		return ArticleModel.Article{}
	}

	// 重新查询以获取完整信息
	var updatedArticle ArticleModel.Article
	result = database.DB.Preload("User").First(&updatedArticle, article.ID)
	if result.Error != nil {
		return ArticleModel.Article{}
	}

	return updatedArticle
}

// Delete 删除文章
func (repository *ArticleRepository) Delete(article ArticleModel.Article) error {

	result := database.DB.Delete(&article)

	// 检查是否删除成功
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到要删除的文章")
	}
	return nil
}

// Search 搜索文章
func (repository *ArticleRepository) Search(keyword string) []ArticleModel.Article {
	var articles []ArticleModel.Article

	// 使用 LIKE 查询搜索标题和内容
	database.DB.Where("title LIKE ? OR content LIKE ?",
		"%"+keyword+"%",
		"%"+keyword+"%").
		Order("created_at DESC").
		Preload("User").
		Find(&articles)

	// 返回搜索结果
	return articles
}
