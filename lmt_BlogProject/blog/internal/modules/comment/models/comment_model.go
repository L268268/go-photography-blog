package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content   string `gorm:"type:text"`
	UserID    uint   `gorm:"not null"`
	ArticleID uint   `gorm:"not null"`
}
