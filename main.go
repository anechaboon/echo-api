package main

import (
	"echo-api/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	InitDB()

	e := echo.New()
	todoHandler := &handlers.TodoHandler{DB: DB}

	e.GET("/todos", todoHandler.GetTodos)
	e.POST("/todos", todoHandler.CreateTodo)
	e.GET("/todos/:id", todoHandler.GetTodoByID)
	e.PUT("/todos/:id", todoHandler.UpdateTodo)
	e.DELETE("/todos/:id", todoHandler.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
