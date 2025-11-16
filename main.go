package main

import (
	"github/AdeleyeShina/initializers"
	"github/AdeleyeShina/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV_variables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	routes.PostRoute(r)

	r.Run()
}
