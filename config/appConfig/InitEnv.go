package appconfig

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lenna-ai/bni-iproc/config"
)

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.DB = config.Oracle()
}
