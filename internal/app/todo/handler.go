package todos

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type TodoHandler struct {
	Storage *TodoStorage
}

func NewTodoHandler(storage *TodoStorage) *TodoHandler {
	return &TodoHandler{Storage: storage}
}

type createTodoRequest struct {
	Title       string
	Description string
}

type updateTodoRequest struct {
	Title       string
	Description string
	Completed   bool
}

type createTodoResponse struct {
	Id int
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	// get the request body
	var body createTodoRequest
	err := c.BodyParser(&body)

	if err != nil {
		return err
	}

	// create the todo
	id, err := h.Storage.CreateNewTodo(NewTodo{
		Title:       body.Title,
		Description: body.Description,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create todo"})
	}

	// send the id
	resp := createTodoResponse{Id: id}

	return c.JSON(resp)
}

func (h *TodoHandler) GetAllTodos(c *fiber.Ctx) error {
	todos, err := h.Storage.GetAllTodos()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not get any todo"})
	}

	return c.JSON(todos)
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	// Get the Todo ID from the URL parameter
	idParam := c.Params("id")

	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing id"})
	}

	// Convert idParam to integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	// Call the delete method on the storage layer
	err = h.Storage.DeleteTodo(id)

	if err != nil {
		// Handle possible errors, like Todo not found
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not delete todo"})
	}

	// Return a success response
	return c.SendStatus(fiber.StatusNoContent)
}

// Update Todo
func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	// Get the Todo ID from the URL parameter
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	// Get the request body
	var body updateTodoRequest
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}

	// Create Todo struct from the request body
	todo := UpdateTodo{
		Id:          id,
		Title:       body.Title,
		Description: body.Description,
		Completed:   body.Completed,
	}

	// Call the storage method to update the Todo
	err = h.Storage.UpdateTodo(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not update todo"})
	}

	return c.SendStatus(fiber.StatusOK)
}

// Complete Todo
// Update Todo
func (h *TodoHandler) CompleteTodo(c *fiber.Ctx) error {
	// Get the Todo ID from the URL parameter
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	// Get the request body
	var body updateTodoRequest
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}

	// Create Todo struct from the request body
	todo := CompletTodo{
		Id:        id,
		Completed: body.Completed,
	}

	// Call the storage method to update the Todo
	err = h.Storage.CompleteTodo(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not update todo"})
	}

	return c.SendStatus(fiber.StatusOK)
}
