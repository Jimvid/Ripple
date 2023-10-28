package static_pages

import (
	"github.com/gofiber/fiber/v2"
)

func home(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {

		features := []string{"Htmx", "Tailwind", "DaisyUi", "Fiber", "Go templates", "Postgres"}

		return c.Render("static_pages/home", fiber.Map{
			"Title":         "Ripple",
			"HeroTitle":     "The Ripple Stack.",
			"HeroParagraph": "The stack that focuses on simplicity, a slim but modern set of tools, and to get you up and running as fast as possible.",
			"HeroFeatures":  features,
		}, "layouts/main")
	})
}
