package api

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hoangduc02011998/todo_server/models"
	"github.com/hoangduc02011998/todo_server/repository/repoIml"
	"github.com/labstack/echo"
)

var userRepo = repoIml.NewUserRepo(models.UserDB.Collection)

func UserPost(c echo.Context) error {
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	var user models.User
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	err = repoIml.NewUserRepo(models.UserDB.Collection).Insert(user)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Invalid,
			Message: err.Error(),
		})
	}

	return Respond(c, &ResponseModel{
		Status:  APIStatus.Ok,
		Message: APIStatus.Ok,
		Data:    user,
	})
}
