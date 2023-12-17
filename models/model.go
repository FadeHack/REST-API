package models

import 	"go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title,omitempty"`
	SubTitle string             `bson:"sub_title,omitempty"`
	Content string             `bson:"content,omitempty"`
}