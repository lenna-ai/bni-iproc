package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/config"
	appconfig "github.com/lenna-ai/bni-iproc/config/appConfig"
	"github.com/lenna-ai/bni-iproc/helpers"
	"github.com/lenna-ai/bni-iproc/routes"
)

func main() {
	defer helpers.RecoverPanicContext(&fiber.Ctx{})
	appconfig.InitApplication()
	app := fiber.New()
	config.Logger(app)
	routes.Router(app)
	if err := app.Listen(":3000"); err != nil {
		panic(err.Error())
	}

}
