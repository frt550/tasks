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
	Db struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	App struct {
		Debug string
	}
}

func init() {
	Config = config{}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// TELEGRAM
	loadEnv("TELEGRAM_API_KEY", func(val string) {
		Config.Telegram.ApiKey = val
	})

	// GRPC
	loadEnv("GRPC_SERVER_NETWORK", func(val string) {
		Config.Grpc.ServerNetwork = val
	})
	loadEnv("GRPC_SERVER_ADDRESS", func(val string) {
		Config.Grpc.ServerAddress = val
	})
	loadEnv("GRPC_CLIENT_TARGET", func(val string) {
		Config.Grpc.ClientTarget = val
	})

	// REST
	loadEnv("REST_SERVER_ADDRESS", func(val string) {
		Config.Rest.ServerAddress = val
	})

	// DB
	loadEnv("DB_HOST", func(val string) {
		Config.Db.Host = val
	})
	loadEnv("DB_PORT", func(val string) {
		Config.Db.Port = val
	})
	loadEnv("DB_USER", func(val string) {
		Config.Db.User = val
	})
	loadEnv("DB_PASSWORD", func(val string) {
		Config.Db.Password = val
	})
	loadEnv("DB_NAME", func(val string) {
		Config.Db.Name = val
	})

	// APP
	loadEnv("APP_DEBUG", func(val string) {
		Config.App.Debug = val
	})
}

func loadEnv(env string, f func(string)) {
	if val, ok := os.LookupEnv(env); ok {
		f(val)
	} else {
		log.Fatalf("Please specify %s in .env file", env)
	}
}
