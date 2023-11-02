package counter

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Router(app *fiber.App) {
	counter := New()
	app.Get("/counter", func(c *fiber.Ctx) error {

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(strconv.Itoa(counter.count))
	})

	app.Post("/counter/increment", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(strconv.Itoa(counter.increment()))
	})

	app.Post("/counter/decrement", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(strconv.Itoa(counter.decrement()))
	})
}
