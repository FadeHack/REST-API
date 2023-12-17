package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FadeHack/Rest-API/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewBlog(t *testing.T) {
	// Create a new HTTP request
	item := models.Item{Title: "Test Title", SubTitle: "Test SubTitle", Content: "Test Content"}
	body, _ := json.Marshal(item)
	req, err := http.NewRequest("POST", "/blog", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	router := gin.Default()
	router.POST("/blog", CreateNewBlog)

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var responseItem models.Item
	err = json.Unmarshal(rr.Body.Bytes(), &responseItem)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, item.Title, responseItem.Title)
	assert.Equal(t, item.SubTitle, responseItem.SubTitle)
	assert.Equal(t, item.Content, responseItem.Content)
}

func TestGetAllBlogs(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/blogs", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	router := gin.Default()
	router.GET("/blogs", GetAllBlogs)

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var responseItems []models.Item
	err = json.Unmarshal(rr.Body.Bytes(), &responseItems)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetBlogById(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/blog/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	router := gin.Default()
	router.GET("/blog/:id", GetBlogById)

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var responseItem models.Item
	err = json.Unmarshal(rr.Body.Bytes(), &responseItem)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUpdateBlogById(t *testing.T) {
	// Create a new HTTP request
	item := models.Item{Title: "Updated Title", SubTitle: "Updated SubTitle", Content: "Updated Content"}
	body, _ := json.Marshal(item)
	req, err := http.NewRequest("PUT", "/ublog/123", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	router := gin.Default()
	router.PUT("/ublog/:id", UpdateBlogById)

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var responseItem models.Item
	err = json.Unmarshal(rr.Body.Bytes(), &responseItem)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, item.Title, responseItem.Title)
	assert.Equal(t, item.SubTitle, responseItem.SubTitle)
	assert.Equal(t, item.Content, responseItem.Content)
}

func TestDeleteBlogById(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("DELETE", "/dblog/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	router := gin.Default()
	router.DELETE("/dblog/:id", DeleteBlogById)

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var responseItem models.Item
	err = json.Unmarshal(rr.Body.Bytes(), &responseItem)
	if err != nil {
		t.Fatal(err)
	}

}
