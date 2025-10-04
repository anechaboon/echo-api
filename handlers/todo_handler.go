package handlers

import (
	"echo-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TodoHandler struct {
	DB *gorm.DB
}

// GET /todos
func (h *TodoHandler) GetTodos(c echo.Context) error {
	var todos []models.Todo
	h.DB.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

// POST /todos
func (h *TodoHandler) CreateTodo(c echo.Context) error {
	println("log:echo-api:Creating a new todo")
	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	h.DB.Create(todo)
	return c.JSON(http.StatusCreated, todo)
}

// GET /todos/:id
func (h *TodoHandler) GetTodoByID(c echo.Context) error {
	id := c.Param("id")
	println("log:echo-api:Fetching todo by ID ", id)
	var todo models.Todo
	if err := h.DB.First(&todo, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
	}
	return c.JSON(http.StatusOK, todo)
}

type UpdateTodoRequest struct {
	Title     *string `json:"title"` // ใช้ pointer จะได้เช็คว่า field ถูกส่งมาหรือไม่
	Completed *bool   `json:"completed"`
}

// PUT /todos/:id
func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	println("log:echo-api:Updating todo by ID ", id)

	var todo models.Todo
	if err := h.DB.First(&todo, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
	}

	var req UpdateTodoRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Update fields ถ้ามีข้อมูลส่งมา
	if req.Title != nil {
		todo.Title = *req.Title
	}
	if req.Completed != nil {
		todo.Completed = *req.Completed
	}

	h.DB.Save(&todo)

	return c.JSON(http.StatusOK, todo)
}

// DELETE /todos/:id
func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	println("log:echo-api:Deleting todo by ID ", id)
	var todo models.Todo
	if err := h.DB.First(&todo, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
	}

	h.DB.Delete(&todo)
	return c.NoContent(http.StatusNoContent)
}
