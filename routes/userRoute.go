package routes

import (
	"github/AdeleyeShina/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	userRoute := r.Group("/api/auth")

	userRoute.POST("/signup", controllers.Signup)
	userRoute.POST("/login", controllers.Login)
	userRoute.POST("/logout", controllers.Logout)
}
