package main

import (
	"github.com/jimvid/ripple/internal/router"
	"log"
)

func main() {
	app := router.New()

	app.Static("/", "./public")

	log.Fatal(app.Listen(":3000"))
}
