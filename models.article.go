package main

import (
	"errors"
)

// Article defines the domain structure supporting both JSON and XML tags
type Article struct {
	ID      int    `json:"id" xml:"id"`
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

// In-memory data store for the microservice
var articleList = []Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// getAllArticles returns all available articles
func getAllArticles() []Article {
	return articleList
}

// getArticleByID fetches a single article by ID or returns an error
func getArticleByID(id int) (*Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}