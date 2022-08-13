package add

import (
	"context"
	"fmt"
	commandPkg "tasks/internal/pkg/bot/command"
	taskPkg "tasks/internal/pkg/core/task"
	taskErr "tasks/internal/pkg/core/task/error"
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
		return taskErr.Error(err)
	} else {
		return fmt.Sprintf("Task %v is added", task)
	}
}
