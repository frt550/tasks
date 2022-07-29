package storage

import (
	"fmt"
	"time"
)

var lastId = uint(0)

type Task struct {
	id          uint
	title       string
	isCompleted bool
	createdAt   time.Time
	completedAt time.Time
}

func (t Task) String() string {
	return fmt.Sprintf("#%d: %s", t.id, t.title)
}

func (t Task) GetId() uint {
	return t.id
}

func (t Task) GetTitle() string {
	return t.title
}

func (t Task) IsCompleted() bool {
	return t.isCompleted
}

func (t Task) GetCreatedAt() time.Time {
	return t.createdAt
}

func (t Task) GetCompletedAt() time.Time {
	return t.completedAt
}

func (t *Task) SetCompletedAt(completedAt time.Time) {
	t.completedAt = completedAt
}
