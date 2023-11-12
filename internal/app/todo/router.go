package todos

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Router(app *fiber.App, conn *sqlx.DB) {
	storage := NewTodoStorage(conn)
	todoHandlers := NewTodoHandler(storage)

	todoGroup := app.Group("/todos")
	todoGroup.Get("/", todoHandlers.GetAllTodos)
	todoGroup.Post("/", todoHandlers.CreateTodo)
	todoGroup.Put("/:id", todoHandlers.UpdateTodo)
	todoGroup.Delete("/:id", todoHandlers.DeleteTodo)
	todoGroup.Put("/:id/completed", todoHandlers.CompleteTodo)
}
