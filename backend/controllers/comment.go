package controllers

import (
	"backend/global"
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(ctx *gin.Context) {
	var comment models.Comment

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.AutoMigrate(&comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.Create(&comment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

func GetComments(ctx *gin.Context) {
	articleId := ctx.Param("article_id")

	var comments []models.Comment

	if err := global.Db.Where("article_id = ?", articleId).Find(&comments).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
