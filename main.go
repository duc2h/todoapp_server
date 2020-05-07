package main

import (
	"github.com/hoangduc02011998/todo_server/action"
	"github.com/hoangduc02011998/todo_server/driver"
	"github.com/hoangduc02011998/todo_server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Response().Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 		c.Header().Set("Access-Control-Allow-Origin", "*")
// w.Header().Set("Access-Control-Allow-Methods", "PUT")
// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 		return next(c)
// 	}
// }

func main() {
	// Connect to DB
	mongo := driver.ConnectMongoDB()
	// Create column todo
	models.InitToDoDB(mongo.Client)

	// api
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://10.17.8.237:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/api/task/all", action.ToDoGetAll)
	e.GET("/api/task", action.ToDoGetByTask)
	e.POST("/api/task", action.ToDoPost)
	e.PUT("/api/task", action.ToDoPut)
	e.DELETE("/api/task", action.ToDoDelete)

	e.Start(":8080")
}
