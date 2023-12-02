package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/jimvid/ripple/config"
	"github.com/jimvid/ripple/internal/app/counter"
	"github.com/jimvid/ripple/internal/app/static_pages"
	"github.com/jimvid/ripple/internal/database"

	todos "github.com/jimvid/ripple/internal/app/todo"
	seed "github.com/jimvid/ripple/internal/database/seeds"
)

func main() {

	engine := html.New("./internal/views", ".html")

	env, err := config.LoadEnv()

	if err != nil {
		log.Fatal(err)
	}

	db := database.CreateMySqlConnection(env)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Images & favicons etc
	app.Static("/", "./assets")

	// non built css, will be removed when we have a build step
	app.Static("/", "./static")

	// Bundled CSS & JS files
	app.Static("/", "./dist")

	if env.ENVIRONMENT != "production" {
		app.Get("/seed", func(c *fiber.Ctx) error {
			seed := seed.New(db)
			seed.SeedTodo()
			return c.SendString("Seeding was completed")
		})
	}

	// Static pages
	static_pages.Router(app)

	// Routes
	todos.Router(app, db)
	counter.Router(app)

	log.Fatal(app.Listen(":3000"))
}
