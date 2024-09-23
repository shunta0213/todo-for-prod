package main

import (
	"todo/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// TasksCreate - Create a new task
	e.POST("/todoapp/task", c.TasksCreate)

	// TasksDelete - Delete an existing task
	e.DELETE("/todoapp/task/:taskId", c.TasksDelete)

	// TasksGetAll - Get the list of all tasks
	e.GET("/todoapp/task", c.TasksGetAll)

	// TasksRead - Get a single task based on its id
	e.GET("/todoapp/task/:taskId", c.TasksRead)

	// TasksUpdate - Update an existing task
	e.PUT("/todoapp/task/:taskId", c.TasksUpdate)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
