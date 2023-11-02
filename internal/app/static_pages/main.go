package static_pages

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	home(app)
	about(app)
}
