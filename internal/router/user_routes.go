package router

import (
	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(app *fiber.App) {
	userGroup := app.Group("/user")

	userGroup.Get("/:id", getUser)
	userGroup.Post("/", createUser)
	// ... add other user routes here
}

func getUser(c *fiber.Ctx) error {
	// Example handler function
	return c.SendString("Get a user!")
}

func createUser(c *fiber.Ctx) error {
	// Example handler function
	return c.SendString("Create a new user!")
}
