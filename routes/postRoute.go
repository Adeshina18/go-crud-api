package routes

import (
	"github/AdeleyeShina/controllers"
	"github/AdeleyeShina/middleware"

	"github.com/gin-gonic/gin"
)

func PostRoute(r *gin.Engine) {
	postRoute := r.Group("/api/post")

	postRoute.GET("/", controllers.GellAllPost)
	postRoute.GET("/:id", controllers.GellSinglePost)

	protectedRoute := postRoute
	protectedRoute.Use(middleware.AuthMiddleWare)
	protectedRoute.POST("/", controllers.CreatePost)
	protectedRoute.PUT("/:id", controllers.UpdatePost)
	protectedRoute.DELETE("/:id", controllers.DeletePost)
}
