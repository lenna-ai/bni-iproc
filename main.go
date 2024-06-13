package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/config"
	appconfig "github.com/lenna-ai/bni-iproc/config/appConfig"
	"github.com/lenna-ai/bni-iproc/routes"
)

func main() {
	appconfig.InitApplication()
	app := fiber.New()
	config.Logger(app)
	routes.Router(app)
	app.Listen(":3000")
}
