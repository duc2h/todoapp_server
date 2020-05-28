package main

import (
	"github.com/hoangduc02011998/todo_server/api"
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
	api.InitUserRepo()
	api.InitTaskRepo()

	// api
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	jwtGroup := e.Group("")
	jwtGroup.Use(isLoggedIn)
	// Task
	jwtGroup.GET("/api/task/all", api.TaskGetAll)
	jwtGroup.GET("/api/task", api.TaskGetByTask)
	jwtGroup.POST("/api/task", api.TaskPost)
	jwtGroup.PUT("/api/task", api.TaskPut)
	jwtGroup.DELETE("/api/task", api.TaskDelete)

	// User
	e.POST("/api/user", api.UserPost)
	e.POST("/api/user/login", api.UserLogin)

	e.Start(":8080")
}
