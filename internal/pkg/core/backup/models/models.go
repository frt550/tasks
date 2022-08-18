package models

import (
	"time"
)

type Backup struct {
	Id        uint      `postgres:"id"`
	Data      string    `postgres:"data"`
	CreatedAt time.Time `postgres:"created_at"`
}
