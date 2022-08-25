package read

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
	return "read"
}

func (c *command) Description() string {
	return "prints detailed info about the task, params: <id>"
}

func (c *command) Process(args string) string {
	ctx := context.Background()
	id, err := strconv.Atoi(args)
	if err != nil {
		return "Please, enter a valid task id"
	}

	task, err := c.task.Get(ctx, uint64(id))
	if err != nil {
		return errPkg.Error(err)
	}

	isCompleted := ""
	if task.IsCompleted {
		isCompleted = "It is completed at " + task.CompletedAt
	} else {
		isCompleted = "It is not completed yet"
	}

	return fmt.Sprintf(
		"You are viewing detailed info of task\n%v.\nThis task was created at %s.\n%s",
		task.String(),
		task.CreatedAt,
		isCompleted,
	)
}
