package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var ApiKey string

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	var ok bool
	if ApiKey, ok = os.LookupEnv("TELEGRAM_API_KEY"); !ok {
		log.Fatal("Please specify TELEGRAM_API_KEY in .env file")
	}
}
