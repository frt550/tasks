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
	Task struct {
		Grpc struct {
			ServerNetwork string
			ServerAddress string
			ClientTarget  string
		}
		Rest struct {
			ServerAddress string
		}
	}
	Backup struct {
		Grpc struct {
			ServerNetwork string
			ServerAddress string
			ClientTarget  string
		}
		Rest struct {
			ServerAddress string
		}
		Metric struct {
			HttpAddress string
		}
	}
	Db struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	TestDb struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	App struct {
		Debug string
	}
	Graylog struct {
		Gelf struct {
			Address string
		}
	}
	Kafka struct {
		Broker0 string
	}
}

func init() {
	Config = config{}

	if err := godotenv.Load(); err != nil {
		// For local tests: running test location is different from project root, so define envPath
		envPath := os.ExpandEnv("$GOPATH/src/tasks/.env")
		if err = godotenv.Load(envPath); err != nil {
			// For ci tests
			envPath := os.ExpandEnv("/builds/frt550/tasks/.env")
			if err = godotenv.Load(envPath); err != nil {
				panic(err)
			}
		}
	}

	// TELEGRAM
	loadEnv("TELEGRAM_API_KEY", func(val string) {
		Config.Telegram.ApiKey = val
	})

	// TASK
	loadEnv("TASK_GRPC_SERVER_NETWORK", func(val string) {
		Config.Task.Grpc.ServerNetwork = val
	})
	loadEnv("TASK_GRPC_SERVER_ADDRESS", func(val string) {
		Config.Task.Grpc.ServerAddress = val
	})
	loadEnv("TASK_GRPC_CLIENT_TARGET", func(val string) {
		Config.Task.Grpc.ClientTarget = val
	})
	loadEnv("TASK_REST_SERVER_ADDRESS", func(val string) {
		Config.Task.Rest.ServerAddress = val
	})

	// BACKUP
	loadEnv("BACKUP_GRPC_SERVER_NETWORK", func(val string) {
		Config.Backup.Grpc.ServerNetwork = val
	})
	loadEnv("BACKUP_GRPC_SERVER_ADDRESS", func(val string) {
		Config.Backup.Grpc.ServerAddress = val
	})
	loadEnv("BACKUP_GRPC_CLIENT_TARGET", func(val string) {
		Config.Backup.Grpc.ClientTarget = val
	})
	loadEnv("BACKUP_REST_SERVER_ADDRESS", func(val string) {
		Config.Backup.Rest.ServerAddress = val
	})
	loadEnv("BACKUP_METRIC_HTTP_ADDRESS", func(val string) {
		Config.Backup.Metric.HttpAddress = val
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

	// TEST_DB
	loadEnv("TEST_DB_HOST", func(val string) {
		Config.TestDb.Host = val
	})
	loadEnv("TEST_DB_PORT", func(val string) {
		Config.TestDb.Port = val
	})
	loadEnv("TEST_DB_USER", func(val string) {
		Config.TestDb.User = val
	})
	loadEnv("TEST_DB_PASSWORD", func(val string) {
		Config.TestDb.Password = val
	})
	loadEnv("TEST_DB_NAME", func(val string) {
		Config.TestDb.Name = val
	})

	// APP
	loadEnv("APP_DEBUG", func(val string) {
		Config.App.Debug = val
	})

	// GRAYLOG
	loadEnv("GRAYLOG_GELF_ADDRESS", func(val string) {
		Config.Graylog.Gelf.Address = val
	})

	// KAFKA
	loadEnv("KAFKA_BROKER0", func(val string) {
		Config.Kafka.Broker0 = val
	})
}

func loadEnv(env string, f func(string)) {
	if val, ok := os.LookupEnv(env); ok {
		f(val)
	} else {
		log.Fatalf("Please specify %s in .env file", env)
	}
}
