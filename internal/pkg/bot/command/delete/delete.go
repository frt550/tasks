package delete

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strconv"
	commandPkg "tasks/internal/pkg/bot/command"
	taskPkg "tasks/internal/pkg/core/task"
	"tasks/internal/pkg/core/task/cache/local"
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
	id, err := strconv.Atoi(args)
	if err != nil {
		return "Please, enter a valid task id"
	}

	task, err := c.task.Get(uint(id))
	if errors.Is(err, local.ErrTaskNotExists) {
		return fmt.Sprintf("Sorry, task #%d is not found", id)
	}

	err = c.task.Delete(task.Id)
	if err != nil {
		log.Println(err)
		return "Internal error"
	}

	return fmt.Sprintf("Task %v is deleted", task)
}
