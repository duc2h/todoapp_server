package action

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hoangduc02011998/todo_server/models"
	"github.com/hoangduc02011998/todo_server/repository/repoIml"
	"github.com/labstack/echo"
)

func ToDoGetAll(c echo.Context) error {
	// c.Request().Header.Set("Context-Type", "application/x-www-form-urlencoded")
	// c.Request().Header.Set("Access-Control-Allow-Origin", "")
	todos, err := repoIml.NewToDoRepo(models.ToDoDB.Collection).FindAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, todos)
}

func ToDoGetByTask(c echo.Context) error {
	task := c.QueryParam("task")
	if task == "" {
		return c.String(http.StatusInternalServerError, "")
	}

	todo, err := repoIml.NewToDoRepo(models.ToDoDB.Collection).FindToDoByTask(task)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, todo)
}

// Create one todo
func ToDoPost(c echo.Context) error {

	var todo models.ToDo
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(bodyBytes, &todo)
	if err != nil {
		c.String(http.StatusInternalServerError, "")
	}

	err = repoIml.NewToDoRepo(models.ToDoDB.Collection).Insert(todo)
	if err != nil {
		c.String(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, todo)
}

func ToDoPut(c echo.Context) error {
	var idStr = c.QueryParam("id")
	var statusStr = c.QueryParam("status")
	if idStr == "" || statusStr == "" {
		return c.String(http.StatusInternalServerError, "")
	}

	status := statusStr == "true"

	err := repoIml.NewToDoRepo(models.ToDoDB.Collection).Update(idStr, status)

	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "Update successfully")
}

func ToDoDelete(c echo.Context) error {
	var idStr = c.QueryParam("id")
	if idStr == "" {
		return c.String(http.StatusInternalServerError, "")
	}

	err := repoIml.NewToDoRepo(models.ToDoDB.Collection).Delete(idStr)
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "Delete successfully")
}
