package repoIml

import (
	"context"
	"errors"

	"github.com/hoangduc02011998/todo_server/models"
	repo "github.com/hoangduc02011998/todo_server/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ToDoRepoIml struct {
	Db *mongo.Database
}

func NewToDoRepo(db *mongo.Database) repo.ToDoRepo {
	return &ToDoRepoIml{
		Db: db,
	}
}

func (mongo *ToDoRepoIml) FindToDoByTask(task string) (models.ToDo, error) {

	todo := models.ToDo{}

	result := mongo.Db.Collection("todo").FindOne(context.Background(), bson.M{"task": task})

	err := result.Decode(&todo)
	if err != nil {
		return todo, errors.New("Not found")
	}

	return todo, nil
}

func (mongo *ToDoRepoIml) Insert(model models.ToDo) error {

	bbytes, err := bson.Marshal(model)

	if err != nil {
		return err
	}

	_, err = mongo.Db.Collection("").InsertOne(context.Background(), bbytes)

	if err != nil {
		return err
	}

	return nil
}

func (mongo *ToDoRepoIml) Update(model models.ToDo) error {
	return nil
}

func (mongo *ToDoRepoIml) Delete(id string) error {
	return nil
}
