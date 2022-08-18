package add

import (
	"context"
	"fmt"
	commandPkg "tasks/internal/pkg/bot/command"
	errPkg "tasks/internal/pkg/core/error"
	taskPkg "tasks/internal/pkg/core/task"
)

func New(task taskPkg.Interface) commandPkg.Interface {
	return &command{
		task: task,
	}
}

type command struct {
	task taskPkg.Interface
}

func (c *command) Name() string {
	return "add"
}

func (c *command) Description() string {
	return "adds a new task, params: <title>"
}

func (c *command) Process(args string) string {
	ctx := context.Background()

	if task, err := c.task.Create(ctx, args); err != nil {
		return errPkg.Error(err)
	} else {
		return fmt.Sprintf("Task %v is added", task)
	}
}
