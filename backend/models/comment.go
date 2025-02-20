package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ArticleID uint   `json:"article_id" binding:"required"` // Foreign key linking to Article
	UserID    uint   `json:"user_id" binding:"required"`    // Foreign key linking to User
	Content   string `json:"content" binding:"required"`    // The content of the comment
	Likes     int    `json:"likes" gorm:"default:0"`        // Optional: Likes on the comment
}
