
package static_pages 

import (
	"github.com/gofiber/fiber/v2"
)

func about(app *fiber.App) {
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("static_pages/about", fiber.Map{
			"Title":   "About",
			"Heading": "About something",
		}, "layouts/main")
	})
}
