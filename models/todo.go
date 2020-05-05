package models

import (
	"github.com/hoangduc02011998/todo_server/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ToDo struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty" bson:"task,omitempty"`
	Status *bool              `json:"status,omitempty" bson:"status,omitempty"`
}

type DBModel struct {
	Collection *mongo.Collection
}

var ToDoDB = &DBModel{}

func InitToDoDB(client *mongo.Client) *mongo.Collection {
	collection := client.Database(config.DB_NAME).Collection("todo")
	ToDoDB.Collection = collection
	return collection
}
