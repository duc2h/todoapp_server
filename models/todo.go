package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDo struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty" bson:"task,omitempty"`
	Status bool               `json:"status,omitempty" bson:"status,omitempty"`
}
