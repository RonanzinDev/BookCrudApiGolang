package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ronanzindev/book-api-golang/controllers"
	"github.com/ronanzindev/book-api-golang/server/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1") // localhost:5000/api/v1
	{
		users := main.Group("users")
		{
			users.POST("/", controllers.CreateUser)

		}
		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}

		main.POST("login", controllers.Login)
	}
	return router
}
