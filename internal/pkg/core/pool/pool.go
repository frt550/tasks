package pool

import (
	"context"
	"fmt"
	"log"
	"sync"
	"tasks/internal/config"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"

	"github.com/jackc/pgx/v4/pgxpool"
)

var instance *pgxpool.Pool
var once sync.Once

type Interface interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func GetInstance() *pgxpool.Pool {
	once.Do(func() {
		instance = createSingleton()
		fmt.Println("created singleton of pool")
	})
	return instance
}

func createSingleton() *pgxpool.Pool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// connection string
	psqlConn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Config.Db.Host,
		config.Config.Db.Port,
		config.Config.Db.User,
		config.Config.Db.Password,
		config.Config.Db.Name,
	)

	// connect to database
	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("ping database error", err)
	}

	return pool
}
