package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/FadeHack/Rest-API/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/gofr"
)

func setup() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://Temp_User:9BH1EM6p6LWStCxt@mongodatabase.ytbk03l.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(ctx, clientOptions)
	return client, err
}

func wrapGofrHandler(h func(c *gofr.Context) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &gofr.Context{
			Request: func() *http.Request {
				return r
			},
			Response: func() http.ResponseWriter {
				return w
			},
		}
		_, err := h(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func TestCreateNewBlog(t *testing.T) {
	client, err := setup()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	item := models.Item{
		Title: "Test Title",
		SubTitle: "Test SubTitle",
		Content: "Test Content",
	}

	jsonItem, _ := json.Marshal(item)
	req, err := http.NewRequest("POST", "/item", bytes.NewBuffer(jsonItem))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := wrapGofrHandler(func(c *gofr.Context) (interface{}, error) {
		return CreateNewBlog(c, client)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// Test for GetAllBlogs
func TestGetAllBlogs(t *testing.T) {
	client, err := setup()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	req, err := http.NewRequest("GET", "/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := wrapGofrHandler(func(c *gofr.Context) (interface{}, error) {
		return GetAllBlogs(c, client)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// Test for GetBlogById
func TestGetBlogById(t *testing.T) {
	client, err := setup()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	req, err := http.NewRequest("GET", "/item/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := wrapGofrHandler(func(c *gofr.Context) (interface{}, error) {
		return GetBlogById(c, client)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// Test for UpdateBlogById
func TestUpdateBlogById(t *testing.T) {
	client, err := setup()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	item := models.Item{
		Title: "Updated Test Title",
		SubTitle: "Updated Test SubTitle",
		Content: "Updated Test Content",
	}

	jsonItem, _ := json.Marshal(item)
	req, err := http.NewRequest("PUT", "/uitem/{id}", bytes.NewBuffer(jsonItem))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := wrapGofrHandler(func(c *gofr.Context) (interface{}, error) {
		return UpdateBlogById(c, client)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// Test for DeleteBlogById
func TestDeleteBlogById(t *testing.T) {
	client, err := setup()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	req, err := http.NewRequest("DELETE", "/ditem/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := wrapGofrHandler(func(c *gofr.Context) (interface{}, error) {
		return DeleteBlogById(c, client)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
