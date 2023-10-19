package router

import (
	"github.com/gofiber/fiber/v2"
)

func setupHomePage(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title":   "hi!",
			"Heading": "Now we have a new heading",
			"Content": "Content",
		}, "layouts/main")
	})
}
