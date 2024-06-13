package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger(app *fiber.App) {
	createDirStorageLogs()
	file, err := os.OpenFile("./storage/logs/"+time.Now().Format("01-02-2006")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	app.Use(logger.New(logger.Config{
		Output: file,
		Format: "body : ${body} | queryParams : ${queryParams} | reqHeaders : ${reqHeaders} | ${time} | ${status} | ${latency} | ${ip} | ${method} | url : ${url} | path : ${path} | route : ${route} | ${error}\n",
	}))
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func createDirStorageLogs() {
	dir := "./storage/logs"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			fmt.Println(dir, "can't created directory")
		}
		fmt.Println("success created directory", dir)
	} else {
		fmt.Println("The provided directory named", dir, "exists")
	}
}
