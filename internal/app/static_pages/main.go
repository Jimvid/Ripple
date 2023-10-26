package static_pages

import ("github.com/gofiber/fiber/v2")

func New(app *fiber.App) {
	home(app)
	about(app)
}
