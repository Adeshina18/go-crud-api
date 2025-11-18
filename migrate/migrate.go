package main

import (
	"github/AdeleyeShina/initializers"
	"github/AdeleyeShina/models"
	"log"
)

func init() {
	initializers.LoadENV_variables()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Post{}, &models.User{})
	if err != nil {
		log.Fatal("Error migrating database")
		return
	}
}
