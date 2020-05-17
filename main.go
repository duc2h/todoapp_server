package main

import (
	"github.com/hoangduc02011998/todo_server/action"
	"github.com/hoangduc02011998/todo_server/driver"
	"github.com/hoangduc02011998/todo_server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Connect to DB
	mongo := driver.ConnectMongoDB()
	// Create column todo
	models.InitToDoDB(mongo.Client)

	// api
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/api/task/all", action.ToDoGetAll)
	e.GET("/api/task", action.ToDoGetByTask)
	e.POST("/api/task", action.ToDoPost)
	e.PUT("/api/task", action.ToDoPut)
	e.DELETE("/api/task", action.ToDoDelete)

	e.Start(":8080")
}
