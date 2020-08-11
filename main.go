package main

import (
	"github.com/hoangduc02011998/todo_server/api"
	biz "github.com/hoangduc02011998/todo_server/business"
	"github.com/hoangduc02011998/todo_server/driver"
	"github.com/hoangduc02011998/todo_server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var isLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningMethod: "HS512",
	SigningKey:    []byte("mySecret"),
})

func main() {
	// Connect to DB
	mongo := driver.ConnectMongoDB()
	// Create column todo
	models.InitTaskDB(mongo.Client)
	models.InitUserDB(mongo.Client)

	// init repo
	biz.InitRepo()

	// api
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	jwtGroup := e.Group("/todo")
	jwtGroup.Use(isLoggedIn)
	// Task
	jwtGroup.GET("/tasks", api.TaskAllGet)
	jwtGroup.GET("/tasks/:name", api.TaskByNameGet)
	jwtGroup.POST("/tasks", api.TaskPost)
	jwtGroup.PUT("/tasks/:name/:status", api.TaskPut)
	jwtGroup.DELETE("/tasks/:name", api.TaskDelete)

	// User
	e.POST("/todo/users", api.UserPost)
	e.POST("/todo/users/login", api.UserLogin)

	e.Start(":8080")
}
