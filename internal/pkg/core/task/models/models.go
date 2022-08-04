package models

import (
	"fmt"
	"time"
)

type Task struct {
	Id          uint
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

func (t Task) String() string {
	return fmt.Sprintf("#%d: %s", t.Id, t.Title)
}
