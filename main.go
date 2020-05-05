package main

import (
	"github.com/hoangduc02011998/todo_server/action"
	"github.com/hoangduc02011998/todo_server/driver"
	"github.com/hoangduc02011998/todo_server/models"
	"github.com/labstack/echo"
)

func main() {
	// Connect to DB
	mongo := driver.ConnectMongoDB()
	// Create column todo
	models.InitToDoDB(mongo.Client)

	// api
	e := echo.New()
	e.GET("/todo", action.ToDoGet)
	e.POST("/todo", action.ToDoPost)

	e.Start(":3000")
}
