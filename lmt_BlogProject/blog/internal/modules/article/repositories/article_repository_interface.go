package repositories

import (
	ArticleModel "blog/internal/modules/article/models"
)

// ArticleRepositoryInterface 定义文章仓库接口
type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
	Find(id int) ArticleModel.Article
	Create(article ArticleModel.Article) ArticleModel.Article
	Update(article ArticleModel.Article) ArticleModel.Article //更新
	Delete(article ArticleModel.Article) error                //删除
	Search(keyword string) []ArticleModel.Article             //搜索
}
