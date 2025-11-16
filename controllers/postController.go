package controllers

import (
	"errors"
	"github/AdeleyeShina/helper"
	"github/AdeleyeShina/initializers"
	"github/AdeleyeShina/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GellAllPost(ctx *gin.Context) {
	var posts []models.Post

	if err := initializers.DB.Find(&posts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":    "",
			"message": "Internal Server Error",
			"status":  "failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    posts,
		"message": "All Posts",
		"status":  "success",
	})
}

func GellSinglePost(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post

	if !helper.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    "",
			"message": "Invalid id",
			"status":  "failed",
		})
		return
	}

	if err := initializers.DB.First(&post, "id= ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"data":    "",
				"message": "No Post Found",
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

	ctx.JSON(http.StatusOK, gin.H{
		"data":    post,
		"message": "The post ",
		"status":  "success",
	})

}

func CreatePost(ctx *gin.Context) {
	var body models.Post

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    "",
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}

	if err := initializers.DB.Create(&post).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":    "",
			"message": "Internal Server Error",
			"status":  "failed",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    post,
		"message": "Post Created Successfully",
		"status":  "success",
	})
}

func UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post
	var body models.Post

	if !helper.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    "",
			"message": "Invalid Id type",
			"status":  "failed",
		})
		return
	}

	if err := initializers.DB.First(&post, "id= ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			ctx.JSON(http.StatusNotFound, gin.H{
				"data":    "",
				"message": "Post Not found",
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

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    "",
			"message": err.Error(),
			"status":  "failed",
		})
		return
	}

	post.Title = body.Title
	post.Body = body.Body

	if err := initializers.DB.Updates(&post).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":    "",
			"message": "Internal Server Error",
			"status":  "failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    post,
		"message": "Post updated",
		"status":  "success",
	})

}

func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")

	if !helper.IsValidUUID(id) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":    "",
			"message": "Invalid Id type",
			"status":  "failed",
		})
		return
	}

	result := initializers.DB.Delete(&models.Post{}, "id=? ", id)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":    "",
			"message": "Internal Server Error",
			"status":  "failed",
		})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"data":    "",
			"message": "Post not found",
			"status":  "failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    "",
		"message": "Deleted",
		"status":  "sucess",
	})

}
