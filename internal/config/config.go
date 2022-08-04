package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config config

type config struct {
	Telegram struct {
		ApiKey string
	}
	Grpc struct {
		ServerNetwork string
		ServerAddress string
		ClientTarget  string
	}
	Rest struct {
		ServerAddress string
	}
}

func init() {
	Config = config{}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	loadEnv("TELEGRAM_API_KEY", func(val string) {
		Config.Telegram.ApiKey = val
	})
	loadEnv("GRPC_SERVER_NETWORK", func(val string) {
		Config.Grpc.ServerNetwork = val
	})
	loadEnv("GRPC_SERVER_ADDRESS", func(val string) {
		Config.Grpc.ServerAddress = val
	})
	loadEnv("GRPC_CLIENT_TARGET", func(val string) {
		Config.Grpc.ClientTarget = val
	})
	loadEnv("HTTP_SERVER_ADDRESS", func(val string) {
		Config.Rest.ServerAddress = val
	})
}

func loadEnv(env string, f func(string)) {
	if val, ok := os.LookupEnv(env); ok {
		f(val)
	} else {
		log.Fatalf("Please specify %s in .env file", env)
	}
}
