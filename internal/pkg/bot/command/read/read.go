package read

import (
	"fmt"
	"strconv"
	commandPkg "tasks/internal/pkg/bot/command"
	taskPkg "tasks/internal/pkg/core/task"
	"tasks/internal/pkg/core/task/cache/local"
	"time"

	"github.com/pkg/errors"
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
	id, err := strconv.Atoi(args)
	if err != nil {
		return "Please, enter a valid task id"
	}

	task, err := c.task.Get(uint(id))
	if errors.Is(err, local.ErrTaskNotExists) {
		return fmt.Sprintf("Sorry, task #%d is not found", id)
	}

	isCompleted := ""
	if task.IsCompleted {
		isCompleted = "It is completed at " + task.CompletedAt.Format(time.RFC850)
	} else {
		isCompleted = "It is not completed yet"
	}

	return fmt.Sprintf(
		"You are viewing detailed info of task\n%v.\nThis task was created at %s.\n%s",
		task.String(),
		task.CreatedAt.Format(time.RFC850),
		isCompleted,
	)
}
