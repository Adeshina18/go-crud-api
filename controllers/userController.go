package controllers

import (
	"errors"
	"github/AdeleyeShina/helper"
	"github/AdeleyeShina/initializers"
	"github/AdeleyeShina/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(ctx *gin.Context) {
	var body models.User

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    "",
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	if err := helper.ValidateUserInput(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    "",
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":    "",
			"message": "Internal Server Error",
			"status":  "failed",
		})
		return
	}

	userData := models.User{Email: body.Email, Name: body.Name, Password: string(hashedPassword)}
	if err := initializers.DB.Create(&userData).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, gin.H{
				"data":    "",
				"message": "Email already exist",
				"status":  "failed",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":    "",
			"message": "Internal Server error",
			"status":  "failed",
		})
		return
	}
	if err := helper.GenerateTokenAndSetCookies(ctx, userData.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	type UserResponse struct {
		ID    uuid.UUID
		Email string
		Name  string
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": UserResponse{
			ID:    userData.ID,
			Email: userData.Email,
			Name:  userData.Name,
		},
		"message": "User logged in succesfully",
		"status":  "success",
	})

}

func AllUser(ctx *gin.Context) {
	var users []models.User
	type UserResponse struct {
		ID    uuid.UUID
		Name  string
		Email string
	}

	if err := initializers.DB.Find(&users).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":    "",
			"message": "Internal Server Error",
			"status":  "failed",
		})
		return
	}
	var response []UserResponse
	for _, user := range users {
		response = append(response, UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    response,
		"message": "Users list",
		"status":  "success",
	})
}

func Login(ctx *gin.Context) {

	type Login struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body Login

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    "",
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	var user models.User

	if err := initializers.DB.First(&user, "email = ?", body.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"data":    "",
				"message": "Account not found",
				"status":  "failed",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":    "",
			"message": "Internal Server Error",
			"status":  "failed",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    "",
			"message": "Invalid username or password",
			"status":  "failed",
		})
		return
	}

	if err := helper.GenerateTokenAndSetCookies(ctx, user.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	type UserResponse struct {
		ID    uuid.UUID
		Email string
		Name  string
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
		"message": "User logged in succesfully",
		"status":  "success",
	})
}

func Logout(ctx *gin.Context) {
	ctx.SetCookie(
		"accessToken",
		"",
		-1,
		"/",
		"",
		true,
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
