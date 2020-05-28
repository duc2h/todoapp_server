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

type TaskRepoIml struct {
	Collection *mongo.Collection
}

func NewTaskRepo(col *mongo.Collection) repo.TaskRepo {
	return &TaskRepoIml{
		Collection: col,
	}
}

func (mongo *TaskRepoIml) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	cur, err := mongo.Collection.Find(context.Background(), bson.M{})

	if err != nil {
		return tasks, errors.New("Not found")
	}

	err = cur.All(context.Background(), &tasks)
	if err != nil {
		return tasks, errors.New("Not found")
	}

	return tasks, nil
}

func (mongo *TaskRepoIml) FindTaskByTask(taskName string) (models.Task, error) {

	task := models.Task{}
	result := mongo.Collection.FindOne(context.Background(), models.Task{TaskName: taskName})

	err := result.Decode(&task)
	if err != nil {
		return task, errors.New("Not found")
	}

	return task, nil
}

func (mongo *TaskRepoIml) Insert(model models.Task) error {

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

func (mongo *TaskRepoIml) Update(id string, status bool) error {
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

func (mongo *TaskRepoIml) Delete(id string) error {

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
