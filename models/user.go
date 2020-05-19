package models

import (
	"github.com/hoangduc02011998/todo_server/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	UserName        string `json:"userName,omitempty" bson:"userName,omitempty"`
	Password        string `json:"password,omitempty" bson:"password,omitempty"`
	PasswordConfirm string `json:"passwordConfirm,omitempty" bson:"password_confirm,omitempty"`
}

var UserDB = &DBModel{}

func InitUserDB(client *mongo.Client) *mongo.Collection {
	collection := client.Database(config.DB_NAME).Collection("user")
	UserDB.Collection = collection
	return collection
}
