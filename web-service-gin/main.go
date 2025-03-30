package main

import (
	"example/web-service-gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)

	router.Run("localhost:8080")
}
