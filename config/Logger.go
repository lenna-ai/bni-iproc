package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func generalLogFile() *os.File {
	generalLogFile, err := os.OpenFile("./storage/logs/general_log/"+time.Now().Format("01-02-2006")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
    if err != nil {
		log.Fatalf("error opening file: %v", err)
    }
	return generalLogFile
}

func Logger(app *fiber.App) {
	createDirStorageLogs()

	generalLogFile := generalLogFile()

	currentDate := time.Now().Format("01-02-2006")
	file, err := os.OpenFile("./storage/logs/"+time.Now().Format("01-02-2006")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	app.Use(func(c *fiber.Ctx) error {
		// Simpan body asli untuk digunakan nanti
		body := c.Body()
		var bodyJSON interface{}
		if err := json.Unmarshal(body, &bodyJSON); err == nil {
			compactBody, _ := json.Marshal(bodyJSON)
			c.Locals("body", string(compactBody))
		} else {
			c.Locals("body", string(body))
		}
		
		return c.Next()
	})

	app.Use(logger.New(logger.Config{
		Output: io.MultiWriter(file, generalLogFile),
		Format: fmt.Sprintf("body : ${locals:body} | queryParams : ${queryParams} | reqHeaders : ${reqHeaders} | time : ${time} | date : %s | status : ${status} | responseTime : ${latency} | ip : ${ip} | ${method} | url : ${url} | path : ${path} | route : ${route} | error : ${error}\n",currentDate),
	}))
	
	log.SetOutput(generalLogFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func createDirStorageLogs() {
	dir := "./storage/logs/general_log"
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
