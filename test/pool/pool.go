package pool

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"sync"
	"tasks/internal/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

var testPool *pgxpool.Pool
var once sync.Once

func GetInstance() (pgx.Tx, func()) {
	once.Do(func() {
		testPool = createSingleton()
		fmt.Println("created singleton of pool")
	})

	tx, err := testPool.Begin(context.Background())
	if err != nil {
		panic(err)
	}
	cleanup := func() {
		if err := tx.Rollback(context.Background()); err != nil {
			panic(err)
		}
	}
	return tx, cleanup
}

func createSingleton() *pgxpool.Pool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// connection string
	psqlConn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Config.TestDb.Host,
		config.Config.TestDb.Port,
		config.Config.TestDb.User,
		config.Config.TestDb.Password,
		config.Config.TestDb.Name,
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
