package responses

import (
	ArticleModel "blog/internal/modules/article/models"
	UserResponse "blog/internal/modules/user/responses"
	"fmt"
)

// 文章响应结构体
type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      UserResponse.User
	UserID    uint //文章作者ID
}

// 文章列表响应结构体
type Articles struct {
	Data []Article
}

// 将文章模型转换为文章响应结构体
func ToArticle(article ArticleModel.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Image:     fmt.Sprintf("/assets/demopic/%d.jpg", article.ID), //修改原有图片，使图片ID与文章ID对应
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      UserResponse.ToUser(article.User),
		UserID:    article.UserID, //文章作者ID
	}
}

// 将文章模型列表转换为文章列表响应结构体
func ToArticles(articles []ArticleModel.Article) Articles {
	var response Articles

	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}

	return response
}
