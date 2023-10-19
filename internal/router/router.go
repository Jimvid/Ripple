package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// New returns a new router instance
func New() *fiber.App {

	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Home page
	setupHomePage(app)

	// User routes
	setupUserRoutes(app)

	return app
}
