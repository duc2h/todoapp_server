package action

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hoangduc02011998/todo_server/models"
	"github.com/hoangduc02011998/todo_server/repository/repoIml"
	"github.com/labstack/echo"
)

func ToDoGet(c echo.Context) error {
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
