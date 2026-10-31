package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/", showIndexPage)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		bodyBytes, err := io.ReadAll(w.Body) // Modern io.ReadAll replacement
		pageOK := err == nil && strings.Contains(string(bodyBytes), "<title>Home Page</title>")

		return statusOK && pageOK
	})
}

func TestArticleListJSON(t *testing.T) {
	r := getRouter(false)
	r.GET("/", showIndexPage)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		bodyBytes, err := io.ReadAll(w.Body)
		expected := `[{"id":1,"title":"Article 1","content":"Article 1 body"},{"id":2,"title":"Article 2","content":"Article 2 body"}]`

		return statusOK && err == nil && strings.TrimSpace(string(bodyBytes)) == expected
	})
}