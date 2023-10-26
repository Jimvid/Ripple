package static_pages 

import (
	"github.com/gofiber/fiber/v2"
)

func home(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("static_pages/home", fiber.Map{
			"Title":   "Home",
			"Heading": "Now we have a new heading",
		}, "layouts/main")
	})
}
