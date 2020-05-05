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
	Collection *mongo.Collection
}

func NewToDoRepo(col *mongo.Collection) repo.ToDoRepo {
	return &ToDoRepoIml{
		Collection: col,
	}
}

func (mongo *ToDoRepoIml) FindToDoByTask(task string) (models.ToDo, error) {

	todo := models.ToDo{}
	result := mongo.Collection.FindOne(context.Background(), bson.M{"task": task})

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

	_, err = mongo.Collection.InsertOne(context.Background(), bbytes)

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
