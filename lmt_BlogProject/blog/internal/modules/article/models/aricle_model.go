package models

import (
	"blog/internal/modules/user/models"

	"gorm.io/gorm"
)

// Article 文章模型
type Article struct {
	gorm.Model
	Title   string `gorm:"varchar:191"`
	Content string `gorm:"text"`
	UserID  uint
	User    models.User
}
