package main

import (
	"log"
	"context"
	"os"
	"time"

	"github.com/FadeHack/Rest-API/routes"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	app.GET("/itemS", func(c *gofr.Context) (interface{}, error) {
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

// // Create an item
// func createItem(c *gofr.Context) (interface{}, error) {
// 	var item Item
// 	err := json.NewDecoder(c.Request().Body).Decode(&item)
// 	if err != nil {
// 		return nil, err
// 	}
// 	collection := client.Database("blog_gofr").Collection("blog")
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	result, _ := collection.InsertOne(ctx, item)
// 	return result, nil
// }

// // Get all items
// func getItems(c *gofr.Context) (interface{}, error) {
// 	var items []Item
// 	collection := client.Database("blog_gofr").Collection("blog")
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	cursor, err := collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		var item Item
// 		cursor.Decode(&item)
// 		items = append(items, item)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// // Get an item by ID
// func getItem(c *gofr.Context) (interface{}, error) {
// 	id, _ := primitive.ObjectIDFromHex(c.PathParam("id"))
// 	var item Item
// 	collection := client.Database("blog_gofr").Collection("blog")
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	err := collection.FindOne(ctx, Item{ID: id}).Decode(&item)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return item, nil
// }



// // Update an item by ID
// func updateItem(c *gofr.Context) (interface{}, error) {
// 	id, _ := primitive.ObjectIDFromHex(c.PathParam("id"))
// 	var item Item
// 	err := json.NewDecoder(c.Request().Body).Decode(&item)
// 	if err != nil {
// 		return nil, err
// 	}
// 	collection := client.Database("blog_gofr").Collection("blog")
// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	result, err := collection.UpdateOne(ctx, Item{ID: id}, bson.M{"$set": item})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// // Delete an item by ID
// func deleteItem(c *gofr.Context) (interface{}, error) {
//     did, _ := primitive.ObjectIDFromHex(c.PathParam("id"))
//     collection := client.Database("blog_gofr").Collection("blog")
//     ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

//     // Find the document before deleting it
//     var item Item
//     err := collection.FindOne(ctx, bson.M{"_id": did}).Decode(&item)
//     if err != nil {
//         return nil, err
//     }

//     // Delete the document
//     _, err = collection.DeleteOne(ctx, bson.M{"_id": did})
//     if err != nil {
//         return nil, err
//     }

//     return item, nil
// }

// func deleteItem(c *gofr.Context) (interface{}, error) {
// 	var cId = c.PathParam("id")
// 	objID, err := primitive.ObjectIDFromHex(cId)
// 	if err != nil {
// 		return primitive.NilObjectID, err
// 	}
// 	collection := client.Database("blog_gofr").Collection("blog")

// 	filter := bson.M{"caseid": objID}

// 	result, err := collection.DeleteOne(c, filter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result.DeletedCount, nil
// }




