package complete

import (
	"context"
	"fmt"
	"strconv"
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
	return "complete"
}

func (c *command) Description() string {
	return "completes the task, params: <id>"
}

func (c *command) Process(args string) string {
	ctx := context.Background()
	id, err := strconv.Atoi(args)
	if err != nil {
		return "Please, enter a valid task id"
	}

	if task, err := c.task.Complete(ctx, uint(id)); err != nil {
		return taskErr.Error(err)
	} else {
		return fmt.Sprintf("Task %v is completed", task)
	}
}
