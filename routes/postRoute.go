package routes

import (
	"github/AdeleyeShina/controllers"

	"github.com/gin-gonic/gin"
)

func PostRoute(r *gin.Engine) {
	postRoute := r.Group("/api/post")
	{
		postRoute.GET("/", controllers.GellAllPost)
		postRoute.GET("/:id", controllers.GellSinglePost)
		postRoute.POST("/", controllers.CreatePost)
		postRoute.PUT("/:id", controllers.UpdatePost)
		postRoute.DELETE("/:id", controllers.DeletePost)
	}
}
