package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	if len(alist) != len(articleList) {
		t.Fatalf("Expected %d articles, got %d", len(articleList), len(alist))
	}

	for i, v := range alist {
		if v.ID != articleList[i].ID || v.Title != articleList[i].Title || v.Content != articleList[i].Content {
			t.Fatalf("Article mismatch at index %d: expected %+v, got %+v", i, articleList[i], v)
		}
	}
}

func TestGetArticleByID(t *testing.T) {
	// Valid ID
	article, err := getArticleByID(1)
	if err != nil || article == nil {
		t.Fatalf("Expected to find article with ID 1, got error: %v", err)
	}

	// Invalid ID
	_, err = getArticleByID(999)
	if err == nil {
		t.Fatal("Expected error for non-existent article ID 999, got nil")
	}
}