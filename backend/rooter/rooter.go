package rooter

import (
	"backend/controllers"
	"backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRooter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	article := r.Group("/api/article")
	article.Use(middlewares.AuthMiddleware())
	{
		article.GET("/", controllers.GetArticles)
		article.GET("/:id", controllers.GetArticle)
		article.POST("/", controllers.CreateArticle)
		// article.PUT("/:id", controllers.UpdateArticle)
		// article.DELETE("/:id", controllers.DeleteArticle)
	}

	comment := r.Group("/api/comment")
	comment.Use(middlewares.AuthMiddleware())
	{
		comment.GET("/:article_id", controllers.GetComments)
		comment.POST("/", controllers.CreateComment)
		// comment.PUT("/:id", controllers.UpdateComment)
		// comment.DELETE("/:id", controllers.DeleteComment)
	}

	return r
}
