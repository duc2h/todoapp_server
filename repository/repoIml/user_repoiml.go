package repoIml

import (
	"context"

	"github.com/hoangduc02011998/todo_server/models"
	repo "github.com/hoangduc02011998/todo_server/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoIml struct {
	Collection *mongo.Collection
}

func NewUserRepo(col *mongo.Collection) repo.UserRepo {
	return &UserRepoIml{
		Collection: col,
	}
}

func (mongo *UserRepoIml) Insert(model models.User) error {

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

func (mongo *UserRepoIml) Login(model models.User) (*models.User, error) {
	user := models.User{}
	result := mongo.Collection.FindOne(context.Background(), model)

	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (mongo *UserRepoIml) GetByUsername(username string) (*models.User, error) {
	user := models.User{}
	result := mongo.Collection.FindOne(context.Background(), models.User{UserName: username})

	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
