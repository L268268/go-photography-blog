package routes

import (
	"blog/internal/middlewares"
	"blog/internal/modules/comment/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册路由
func RegisterRoutes(router *gin.Engine) {
	controller := controllers.New()

	authGroup := router.Group("/comments")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.POST("", controller.Comment) // 评论
	}
}
