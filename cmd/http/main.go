package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jimvid/ripple/internal/app/counter"
	"github.com/jimvid/ripple/internal/app/static_pages"
)

func main() {

	engine := html.New("./internal/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Images & favicons etc
	app.Static("/", "./assets")

	// non built css, will be removed when we have a build step
	app.Static("/", "./static")

	// Bundled CSS & JS files
	app.Static("/", "./dist")

	// Static pages
	static_pages.Router(app)
	counter.Router(app)

	log.Fatal(app.Listen(":3000"))
}
