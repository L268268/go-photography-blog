package services

import (
	ArticleModel "blog/internal/modules/article/models"
	ArticleRepository "blog/internal/modules/article/repositories"
	"blog/internal/modules/article/requests/articles"
	ArticleResponse "blog/internal/modules/article/responses"
	UserResponse "blog/internal/modules/user/responses"
	"errors"
	"fmt"
)

// 文章服务层
type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

// 获取文章列表
func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

// 获取推荐文章
func (articleService *ArticleService) GetFeaturedArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(4)

	return ArticleResponse.ToArticles(articles)
}

// 获取首页文章
func (articleService *ArticleService) GetStoriesArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(12)

	return ArticleResponse.ToArticles(articles)
}

// 获取所有文章
func (articleService *ArticleService) Find(id int) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article

	article := articleService.articleRepository.Find(id)

	if article.ID == 0 {
		return response, errors.New("article not found")
	}

	return ArticleResponse.ToArticle(article), nil
}

// 创建文章
func (articleService *ArticleService) StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error) {
	var article ArticleModel.Article
	var response ArticleResponse.Article

	article.Title = request.Title
	article.Content = request.Content
	article.UserID = user.ID

	newArticle := articleService.articleRepository.Create(article)

	if newArticle.ID == 0 {
		return response, errors.New("error in creating the article")
	}

	return ArticleResponse.ToArticle(newArticle), nil
}

// 更新文章
func (articleService *ArticleService) Update(article ArticleModel.Article) (ArticleResponse.Article, error) {
	// 首先检查文章是否存在
	existingArticle := articleService.articleRepository.Find(int(article.ID))
	if existingArticle.ID == 0 {
		return ArticleResponse.Article{}, errors.New("article not found")
	}

	// 保留不需要更新的字段
	article.CreatedAt = existingArticle.CreatedAt
	if article.UserID == 0 {
		article.UserID = existingArticle.UserID
	}

	// 更新文章
	updatedArticle := articleService.articleRepository.Update(article)
	if updatedArticle.ID == 0 {
		return ArticleResponse.Article{}, errors.New("error updating article")
	}

	return ArticleResponse.ToArticle(updatedArticle), nil
}

// 删除文章
func (articleService *ArticleService) Delete(id int) error {
	// 首先查找文章是否存在
	article := articleService.articleRepository.Find(id)
	if article.ID == 0 {
		return fmt.Errorf("文章ID %d 不存在", id)
	}

	// 删除文章
	if err := articleService.articleRepository.Delete(article); err != nil {
		return fmt.Errorf("删除文章失败: %v", err)
	}

	return nil
}

// 搜索
func (articleService *ArticleService) Search(keyword string) ArticleResponse.Articles {
	articles := articleService.articleRepository.Search(keyword)
	return ArticleResponse.ToArticles(articles)
}
