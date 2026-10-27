package services

import (
	ArticleModel "blog/internal/modules/article/models"
	"blog/internal/modules/article/requests/articles"
	ArticleResponse "blog/internal/modules/article/responses"
	UserResponse "blog/internal/modules/user/responses"
)

// ArticleServiceInterface 文章服务接口
type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles                                                      //获取推荐文章
	GetStoriesArticles() ArticleResponse.Articles                                                       //获取故事文章
	Find(id int) (ArticleResponse.Article, error)                                                       //获取文章
	StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error) //以用户身份存储文章
	Update(article ArticleModel.Article) (ArticleResponse.Article, error)                               //更新接口
	Delete(id int) error                                                                                //删除接口
	Search(keyword string) ArticleResponse.Articles                                                     //搜索接口
}
