package static_pages

import (
	"github.com/gofiber/fiber/v2"
)

func home(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("static_pages/home", fiber.Map{
			"Title":         "Home",
			"HeroTitle":     "Radix Themes",
			"HeroParagraph": "The open source component library optimized for fast development, easy maintainence, and accessibility. Figma edition.",
		}, "layouts/main")
	})
}
