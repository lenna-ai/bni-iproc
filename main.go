package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lenna-ai/bni-iproc/config"
	appconfig "github.com/lenna-ai/bni-iproc/config/appConfig"
	"github.com/lenna-ai/bni-iproc/helpers"
	"github.com/lenna-ai/bni-iproc/routes"
)

func main() {
	defer helpers.RecoverPanicContext(&fiber.Ctx{})
	appconfig.InitApplication()
	app := fiber.New()
	app.Use(cors.New())
	config.Logger(app)
	routes.Router(app)

	// testLoginController := injector.InitializeController()
	// fmt.Println(testLoginController)

	if err := app.Listen(":3000"); err != nil {
		panic(err.Error())
	}

}
