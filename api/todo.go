package api

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hoangduc02011998/todo_server/models"
	"github.com/hoangduc02011998/todo_server/repository/repoIml"
	"github.com/labstack/echo"
)

func ToDoGetAll(c echo.Context) error {
	todos, err := repoIml.NewToDoRepo(models.ToDoDB.Collection).FindAll()
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Invalid,
			Message: err.Error(),
		})
	}

	return Respond(c, &ResponseModel{
		Status:  APIStatus.Ok,
		Message: APIStatus.Ok,
		Data:    todos,
	})
}

func ToDoGetByTask(c echo.Context) error {
	task := c.QueryParam("task")
	if task == "" {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Invalid,
			Message: APIStatus.Invalid,
		})
	}

	todo, err := repoIml.NewToDoRepo(models.ToDoDB.Collection).FindToDoByTask(task)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	return Respond(c, &ResponseModel{
		Status:  APIStatus.Ok,
		Message: APIStatus.Ok,
		Data:    todo,
	})
}

// Create one todo
func ToDoPost(c echo.Context) error {
	var todo models.ToDo
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	err = json.Unmarshal(bodyBytes, &todo)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	err = repoIml.NewToDoRepo(models.ToDoDB.Collection).Insert(todo)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	return Respond(c, &ResponseModel{
		Status:  APIStatus.Ok,
		Message: APIStatus.Ok,
		Data:    todo,
	})
}

func ToDoPut(c echo.Context) error {
	var idStr = c.QueryParam("id")
	var statusStr = c.QueryParam("status")
	if idStr == "" || statusStr == "" {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Invalid,
			Message: APIStatus.Invalid,
		})
	}

	status := statusStr == "true"
	err := repoIml.NewToDoRepo(models.ToDoDB.Collection).Update(idStr, status)

	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	return Respond(c, &ResponseModel{
		Status:  APIStatus.Ok,
		Message: APIStatus.Ok,
	})
}

func ToDoDelete(c echo.Context) error {
	var idStr = c.QueryParam("id")
	if idStr == "" {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Invalid,
			Message: APIStatus.Invalid,
		})
	}

	err := repoIml.NewToDoRepo(models.ToDoDB.Collection).Delete(idStr)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	return Respond(c, &ResponseModel{
		Status:  APIStatus.Ok,
		Message: APIStatus.Ok,
	})
}
