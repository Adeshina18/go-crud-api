package middleware

import (
	"errors"
	"github/AdeleyeShina/initializers"
	"github/AdeleyeShina/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthMiddleWare(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("accessToken")
	if err != nil || tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Session expired pls login again",
		})
		ctx.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid or expired token",
		})
		ctx.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token claim",
		})
		ctx.Abort()
		return
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token expired",
			})
			ctx.Abort()
			return
		}
	}

	var user models.User
	if err := initializers.DB.First(&user, "id= ?", claims["userId"]).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "User not found",
			})
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		ctx.Abort()
		return
	}
	user.Password = ""
	ctx.Set("user", user)
}
