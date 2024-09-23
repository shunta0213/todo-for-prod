package handlers

import (
	"net/http"
	"todo/models"

	"github.com/labstack/echo/v4"
)

// TasksCreate - Create a new task
func (c *Container) TasksCreate(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// TasksDelete - Delete an existing task
func (c *Container) TasksDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// TasksGetAll - Get the list of all tasks
func (c *Container) TasksGetAll(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// TasksRead - Get a single task based on its id
func (c *Container) TasksRead(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// TasksUpdate - Update an existing task
func (c *Container) TasksUpdate(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
