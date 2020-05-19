package models

import "go.mongodb.org/mongo-driver/mongo"

type DBModel struct {
	Collection *mongo.Collection
}
