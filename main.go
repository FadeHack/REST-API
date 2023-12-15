package main

import (
	"context"
	"crypto/dsa"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/gofr"
)

type Item struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title,omitempty"`
	SubTitle string             `bson:"sub_title,omitempty"`
	Content string             `bson:"content,omitempty"`
}

var client *mongo.Client

func main(){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://Temp_User:9BH1EM6p6LWStCxt@mongodatabase.ytbk03l.mongodb.net/?retryWrites=true&w=majority")
	client, _ = mongo.Connect(ctx, clientOptions)

	app := gofr.New()

	app.POST("/item", createItem)
	app.GET("/items", getItems)
	app.GET("/item/{id}", getItem)
	app.PATCH("/item/{id}", patchItem)
	app.PUT("/item/{id}", updateItem)
	app.DELETE("/item/{id}", deleteItem)

	app.Start()
}
// Create an item
func createItem(c *gofr.Context) (interface{}, error) {
	var item Item
	err := json.NewDecoder(c.Request().Body).Decode(&item)
	if err != nil {
		return nil, err
	}
	collection := client.Database("blog_gofr").Collection("blog")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, item)
	return result, nil
}

// Get all items
func getItems(c *gofr.Context) (interface{}, error) {
	var items []Item
	collection := client.Database("blog_gofr").Collection("blog")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var item Item
		cursor.Decode(&item)
		items = append(items, item)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// Get an item by ID
func getItem(c *gofr.Context) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(c.Params()["id"])
	var item Item
	collection := client.Database("blog_gofr").Collection("blog")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Item{ID: id}).Decode(&item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Update a part of an item by ID
func patchItem(c *gofr.Context) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(c.Params()["id"])
	var update bson.M
	err := json.NewDecoder(c.Request().Body).Decode(&update)
	if err != nil {
		return nil, err
	}
	collection := client.Database("blog_gofr").Collection("blog")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	result, err := collection.UpdateOne(ctx, Item{ID: id}, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	return result, nil
}


// Update an item by ID
func updateItem(c *gofr.Context) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(c.Params()["id"])
	var item Item
	err := json.NewDecoder(c.Request().Body).Decode(&item)
	if err != nil {
		return nil, err
	}
	collection := client.Database("blog_gofr").Collection("blog")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	result, err := collection.UpdateOne(ctx, Item{ID: id}, bson.M{"$set": item})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Delete an item by ID
func deleteItem(c *gofr.Context) (interface{}, error) {
    did, _ := primitive.ObjectIDFromHex(c.PathParam("id"))
    collection := client.Database("blog_gofr").Collection("blog")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

    // Find the document before deleting it
    var item Item
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




