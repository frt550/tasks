package update

import (
	"context"
	"fmt"
	"strconv"
	"strings"
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
	return "update"
}

func (c *command) Description() string {
	return "updates the task, params: <id> <title>"
}

func (c *command) Process(args string) string {
	params := strings.Split(args, " ")
	id, err := strconv.Atoi(params[0])
	if err != nil {
		return "Please, enter a valid task id"
	}
	var title string
	if len(params) >= 2 {
		title = strings.Join(params[1:], " ")
	}

	if task, err := c.task.UpdateTitle(context.Background(), uint(id), title); err != nil {
		return errPkg.Error(err)
	} else {
		return fmt.Sprintf("Task #%d is updated to %s", task.Id, task.Title)
	}
}
