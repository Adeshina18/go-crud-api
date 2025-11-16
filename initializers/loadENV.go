package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadENV_variables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading the ENV Variables")
		return
	}
}
