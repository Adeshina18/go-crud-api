package helper

import (
	"errors"
	"github/AdeleyeShina/models"
	"os"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func ValidateUserInput(user models.User) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	emailRegex := regexp.MustCompile(regex)

	if !emailRegex.MatchString(user.Email) {
		return errors.New("invalid email address")
	}

	if len(user.Password) < 6 {
		return errors.New("password cannot be less than 6 character")
	}
	return nil
}

func GenerateTokenAndSetCookies(ctx *gin.Context, userId uuid.UUID) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	ctx.SetCookie(
		"accessToken",
		tokenString,
		3600*24,
		"/",
		"",
		true,
		true,
	)
	return nil
}
