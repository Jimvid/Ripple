package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jimvid/ripple/internal/app/static_pages"
	"log"
)

func main() {

	engine := html.New("./internal/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

    // Images & favicons etc
    app.Static("/", "./static")

    // Bundled CSS & JS files
    app.Static("/", "./dist")

	// Static pages 
	static_pages.New(app)


	log.Fatal(app.Listen(":3000"))
}
