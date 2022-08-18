package delete

import (
	"context"
	"fmt"
	"strconv"
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
	return "delete"
}

func (c *command) Description() string {
	return "deletes the task, params: <id>"
}

func (c *command) Process(args string) string {
	ctx := context.Background()
	id, err := strconv.Atoi(args)
	if err != nil {
		return "Please, enter a valid task id"
	}

	if task, err := c.task.Delete(ctx, uint(id)); err != nil {
		return errPkg.Error(err)
	} else {
		return fmt.Sprintf("Task %v is deleted", task)
	}
}
