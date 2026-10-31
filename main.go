package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Routes
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)

	return router
}

func main() {
	router := setupRouter()

	// Default listens on 0.0.0.0:8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}