package migration

import (
	articleModels "blog/internal/modules/article/models"
	commentModels "blog/internal/modules/comment/models"
	userModels "blog/internal/modules/user/models"
	"blog/pkg/database"
	"fmt"
	"log"
)

// Migrate 数据库迁移
func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{}, &commentModels.Comment{}) //评论数据库迁移

	// err := db.AutoMigrate(&articleModels.Article{}

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration done ..")
}
