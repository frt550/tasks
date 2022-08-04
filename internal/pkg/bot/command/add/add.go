package add

import (
	"fmt"
	"log"
	"strings"
	commandPkg "tasks/internal/pkg/bot/command"
	"tasks/internal/pkg/core/counter"
	taskPkg "tasks/internal/pkg/core/task"
	"tasks/internal/pkg/core/task/models"
	"time"
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
	var title = strings.TrimSpace(args)
	if len(title) == 0 {
		return "Title cannot be empty"
	}
	var taskModel = models.Task{
		Id:          counter.GetId(),
		Title:       strings.TrimSpace(args),
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}

	if err := c.task.Create(taskModel); err != nil {
		log.Println(err)
		return "Internal error"
	}
	return fmt.Sprintf("Task %v is added", taskModel)
}
