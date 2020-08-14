package api

import (
	"encoding/json"
	"io/ioutil"

	biz "github.com/hoangduc02011998/todo-app/server/business"
	"github.com/hoangduc02011998/todo-app/server/models"
	"github.com/labstack/echo"
)

func TaskAllGet(c echo.Context) error {

	return models.Respond(c, biz.GetAllTask())
}

func TaskByNameGet(c echo.Context) error {
	taskName := c.Param("name")
	if taskName == "" {
		return models.Respond(c, &models.ResponseModel{
			Status:  models.APIStatus.Invalid,
			Message: "Required name of task",
		})
	}

	return models.Respond(c, biz.GetTaskByName(taskName))
}

// Create one Task
func TaskPost(c echo.Context) error {
	var task models.Task
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return models.Respond(c, &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: err.Error(),
		})
	}

	err = json.Unmarshal(bodyBytes, &task)
	if err != nil {
		return models.Respond(c, &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: err.Error(),
		})
	}

	if task.TaskName == "" {
		return models.Respond(c, &models.ResponseModel{
			Status:  models.APIStatus.Invalid,
			Message: "TaskName is required",
		})
	}

	return models.Respond(c, biz.CreateTask(task))
}

func TaskPut(c echo.Context) error {
	var name = c.Param("name")
	var statusStr = c.Param("status")
	if name == "" || statusStr == "" {
		return models.Respond(c, &models.ResponseModel{
			Status:  models.APIStatus.Invalid,
			Message: models.APIStatus.Invalid,
		})
	}

	status := statusStr == "true"
	return models.Respond(c, biz.UpdateTask(name, status))
}

func TaskDelete(c echo.Context) error {
	var name = c.Param("name")
	if name == "" {
		return models.Respond(c, &models.ResponseModel{
			Status:  models.APIStatus.Invalid,
			Message: models.APIStatus.Invalid,
		})
	}

	return models.Respond(c, biz.DeleteTask(name))
}
