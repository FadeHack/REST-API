package main

import (
	"log"
	"context"
	"os"
	"time"

	"github.com/FadeHack/Rest-API/routes"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/gofr"
)

func main(){
	// Load .env file from configs folder
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the URI from the .env file
	uri := os.Getenv("MONGO_URI")

	var client *mongo.Client
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(uri)
	client, _ = mongo.Connect(ctx, clientOptions)

	app := gofr.New()

	app.POST("/item", func(c *gofr.Context) (interface{}, error) {
		return routes.CreateNewBlog(c, client)
	})
	app.GET("/items", func(c *gofr.Context) (interface{}, error) {
		return routes.GetAllBlogs(c, client)
	})
	app.GET("/item/{id}", func(c *gofr.Context) (interface{}, error) {
		return routes.GetBlogById(c, client)
	})
	app.PUT("/uitem/{id}", func(c *gofr.Context) (interface{}, error) {
		return routes.UpdateBlogById(c, client)
	})
	app.DELETE("/ditem/{id}", func(c *gofr.Context) (interface{}, error) {
		return routes.DeleteBlogById(c, client)
	})

	app.Start()
}
