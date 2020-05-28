package api

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hoangduc02011998/todo_server/models"
	repo "github.com/hoangduc02011998/todo_server/repository"
	"github.com/hoangduc02011998/todo_server/repository/repoIml"
	"github.com/labstack/echo"
)

var taskRepo repo.TaskRepo

func InitTaskRepo() {
	taskRepo = repoIml.NewTaskRepo(models.TaskDB.Collection)
}

func TaskGetAll(c echo.Context) error {
	tasks, err := taskRepo.FindAll()
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Invalid,
			Message: err.Error(),
		})
	}

	return Respond(c, &ResponseModel{
		Status:  APIStatus.Ok,
		Message: APIStatus.Ok,
		Data:    tasks,
	})
}

func TaskGetByTask(c echo.Context) error {
	taskName := c.QueryParam("taskName")
	if taskName == "" {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Invalid,
			Message: APIStatus.Invalid,
		})
	}

	task, err := taskRepo.FindTaskByTask(taskName)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	return Respond(c, &ResponseModel{
		Status:  APIStatus.Ok,
		Message: APIStatus.Ok,
		Data:    task,
	})
}

// Create one Task
func TaskPost(c echo.Context) error {
	var task models.Task
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	err = json.Unmarshal(bodyBytes, &task)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	err = taskRepo.Insert(task)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	return Respond(c, &ResponseModel{
		Status:  APIStatus.Ok,
		Message: APIStatus.Ok,
		Data:    task,
	})
}

func TaskPut(c echo.Context) error {
	var idStr = c.QueryParam("id")
	var statusStr = c.QueryParam("status")
	if idStr == "" || statusStr == "" {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Invalid,
			Message: APIStatus.Invalid,
		})
	}

	status := statusStr == "true"
	err := taskRepo.Update(idStr, status)

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

func TaskDelete(c echo.Context) error {
	var idStr = c.QueryParam("id")
	if idStr == "" {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Invalid,
			Message: APIStatus.Invalid,
		})
	}

	err := taskRepo.Delete(idStr)
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
