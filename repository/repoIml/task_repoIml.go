package repoIml

import (
	"context"
	"errors"

	"github.com/hoangduc02011998/todo-app/server/models"
	repo "github.com/hoangduc02011998/todo-app/server/repository"
	"go.mongodb.org/mongo-driver/bson"
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
		return tasks, err
	}

	return tasks, nil
}

func (mongo *TaskRepoIml) FindTaskByName(taskName string) (*models.Task, error) {

	task := models.Task{}
	result := mongo.Collection.FindOne(context.Background(), models.Task{TaskName: taskName})

	err := result.Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
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

func (mongo *TaskRepoIml) Update(name string, status bool) error {
	updater := bson.M{"$set": bson.M{"status": &status}}
	_, err := mongo.Collection.UpdateOne(context.Background(), models.Task{TaskName: name}, updater)
	if err != nil {
		return err
	}

	return nil
}

func (mongo *TaskRepoIml) Delete(name string) error {

	_, err := mongo.Collection.DeleteOne(context.Background(), models.Task{TaskName: name})
	if err != nil {
		return err
	}
	return nil
}
