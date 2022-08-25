package models

import (
	"fmt"
)

type Task struct {
	Id          uint64 `postgres:"id"`
	Title       string `postgres:"title"`
	IsCompleted bool   `postgres:"is_completed"`
	CreatedAt   string `postgres:"created_at"`
	CompletedAt string `postgres:"completed_at"`
}

func (t Task) String() string {
	return fmt.Sprintf("#%d: %s", t.Id, t.Title)
}
