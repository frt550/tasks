package update

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	commandPkg "tasks/internal/pkg/bot/command"
	taskPkg "tasks/internal/pkg/core/task"
	"tasks/internal/pkg/core/task/cache/local"

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
	if len(title) == 0 {
		return "Title cannot be empty"
	}

	task, err := c.task.Get(uint(id))
	if errors.Is(err, local.ErrTaskNotExists) {
		return fmt.Sprintf("Sorry, task #%d is not found", id)
	}
	if task.IsCompleted {
		return "Completed task cannot be updated"
	}

	task.Title = title
	err = c.task.Update(task)
	if err != nil {
		log.Println(err)
		return "Internal error"
	}

	return fmt.Sprintf("Task #%d is updated to %s", task.Id, task.Title)
}
