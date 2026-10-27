package routes

import (
	"blog/internal/middlewares"
	"blog/internal/modules/article/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	controller := controllers.New()

	router.GET("/articles/:id", controller.Show)      //详情路由
	router.GET("/articles/search", controller.Search) //搜索路由

	authGroup := router.Group("/articles") //文章管理路由
	authGroup.Use(middlewares.IsAuth())    //鉴权中间件
	{
		authGroup.GET("/create", controller.Create)      //创建路由
		authGroup.POST("/store", controller.Store)       //创建路由
		authGroup.GET("/:id/edit", controller.Edit)      //编辑路由
		authGroup.POST("/:id/update", controller.Update) //更新路由
		authGroup.DELETE("/:id", controller.Delete)      //删除路由
	}
}
