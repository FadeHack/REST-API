package routes

import (
	"context"
	"encoding/json"
	"time"

	"github.com/FadeHack/Rest-API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gofr.dev/pkg/gofr"
)



func CreateNewBlog(c *gofr.Context, client *mongo.Client) (interface{}, error) {
	var item models.Item
	err := json.NewDecoder(c.Request().Body).Decode(&item)
	if err != nil {
		return nil, err
	}
	collection := client.Database("blog_gofr").Collection("blog")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, item)
	return result, nil
}

func GetAllBlogs(c *gofr.Context, client *mongo.Client) (interface{}, error) {
	var items []models.Item
	collection := client.Database("blog_gofr").Collection("blog")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var item models.Item
		cursor.Decode(&item)
		items = append(items, item)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func GetBlogById(c *gofr.Context, client *mongo.Client) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(c.PathParam("id"))
	var item models.Item
	collection := client.Database("blog_gofr").Collection("blog")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, models.Item{ID: id}).Decode(&item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func UpdateBlogById(c *gofr.Context, client *mongo.Client) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(c.PathParam("id"))
	var item models.Item
	err := json.NewDecoder(c.Request().Body).Decode(&item)
	if err != nil {
		return nil, err
	}
	collection := client.Database("blog_gofr").Collection("blog")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	result, err := collection.UpdateOne(ctx, models.Item{ID: id}, bson.M{"$set": item})
	if err != nil {
		return nil, err
	}
	return result, nil
}


// Delete an item by ID
func DeleteBlogById(c *gofr.Context, client *mongo.Client) (interface{}, error) {
    did, _ := primitive.ObjectIDFromHex(c.PathParam("id"))
    collection := client.Database("blog_gofr").Collection("blog")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    // Find the document before deleting it
    var item models.Item
    err := collection.FindOne(ctx, bson.M{"_id": did}).Decode(&item)
    if err != nil {
        return nil, err
    }

    // Delete the document
    _, err = collection.DeleteOne(ctx, bson.M{"_id": did})
    if err != nil {
        return nil, err
    }

    return item, nil
}