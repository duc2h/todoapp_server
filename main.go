package main

import (
	"github.com/hoangduc02011998/todo_server/api"
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
	models.InitUserDB(mongo.Client)

	// api
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Task
	e.GET("/api/task/all", api.ToDoGetAll)
	e.GET("/api/task", api.ToDoGetByTask)
	e.POST("/api/task", api.ToDoPost)
	e.PUT("/api/task", api.ToDoPut)
	e.DELETE("/api/task", api.ToDoDelete)

	// User
	e.POST("/api/user", api.UserPost)

	e.Start(":8080")
}
