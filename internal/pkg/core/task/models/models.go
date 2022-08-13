package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Task struct {
	Id          uint         `postgres:"id"`
	Title       string       `postgres:"title"`
	IsCompleted bool         `postgres:"is_completed"`
	CreatedAt   time.Time    `postgres:"created_at"`
	CompletedAt sql.NullTime `postgres:"completed_at"`
}

func (t Task) String() string {
	return fmt.Sprintf("#%d: %s", t.Id, t.Title)
}
