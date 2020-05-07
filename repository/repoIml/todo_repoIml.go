package repoIml

import (
	"context"
	"errors"

	"github.com/hoangduc02011998/todo_server/models"
	repo "github.com/hoangduc02011998/todo_server/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (mongo *ToDoRepoIml) FindAll() ([]models.ToDo, error) {
	var todos []models.ToDo
	cur, err := mongo.Collection.Find(context.Background(), bson.M{})

	if err != nil {
		return todos, errors.New("Not found")
	}

	err = cur.All(context.Background(), &todos)
	if err != nil {
		return todos, errors.New("Not found")
	}

	return todos, nil
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

// func (mongo *ToDoRepoIml) CompleteTask(id string) error {
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}

// 	status := true
// 	_, err = mongo.Collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, models.ToDo{Status: &status})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (mongo *ToDoRepoIml) UndoTask(id string) error {
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}

// 	status := false
// 	_, err = mongo.Collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, models.ToDo{Status: &status})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (mongo *ToDoRepoIml) Update(id string, status bool) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	updater := bson.M{"$set": bson.M{"status": &status}}
	_, err = mongo.Collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, updater)
	if err != nil {
		return err
	}

	return nil
}

func (mongo *ToDoRepoIml) Delete(id string) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = mongo.Collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}
