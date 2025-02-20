package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ArticleID uint   `json:"article_id" binding:"required"`
	UserID    uint   `json:"user_id" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Likes     int    `json:"likes" binding:"required" gorm:"default:0"`
}
