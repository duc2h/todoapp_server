package models

import (
	"github.com/hoangduc02011998/todo_server/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	TaskName string             `json:"taskName,omitempty" bson:"task_name,omitempty"`
	Status   *bool              `json:"status,omitempty" bson:"status,omitempty"`
}

var TaskDB = &DBModel{}

func InitTaskDB(client *mongo.Client) *mongo.Collection {
	collection := client.Database(config.DB_NAME).Collection("task")
	TaskDB.Collection = collection
	return collection
}
