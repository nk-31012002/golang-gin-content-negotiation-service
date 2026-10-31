package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// render inspects the Accept header and renders JSON, XML, or HTML
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

// showIndexPage handles requests to the root index
func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

// getArticle handles requests to display an individual article
func getArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	article, err := getArticleByID(articleID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	render(c, gin.H{
		"title":   article.Title,
		"payload": article,
	}, "article.html")
}