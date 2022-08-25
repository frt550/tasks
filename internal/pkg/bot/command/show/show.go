package show

import (
	"context"
	"strings"
	commandPkg "tasks/internal/pkg/bot/command"
	errPkg "tasks/internal/pkg/core/error"
	taskPkg "tasks/internal/pkg/core/task"
	"tasks/internal/pkg/core/task/models"
)

const paramAll = "all"

func New(task taskPkg.Interface) commandPkg.Interface {
	return &command{
		task: task,
	}
}

type command struct {
	task taskPkg.Interface
}

func (c *command) Name() string {
	return "show"
}

func (c *command) Description() string {
	return "shows pending tasks, /show all - shows all tasks"
}

func (c *command) Process(args string) string {
	var allTasks, err = c.task.All(context.Background(), 0, 0)
	if err != nil {
		errPkg.Error(err)
	}
	var pendingTasks = make([]*models.Task, 0, len(allTasks))
	var completedTasks = make([]*models.Task, 0, len(allTasks))

	for _, task := range allTasks {
		if task.IsCompleted {
			completedTasks = append(completedTasks, task)
		} else {
			pendingTasks = append(pendingTasks, task)
		}
	}

	var result []string
	if len(pendingTasks) == 0 {
		result = append(result, "All tasks are done!")
	} else {
		for _, pendingTask := range pendingTasks {
			result = append(result, pendingTask.String())
		}
	}

	params := strings.Split(args, " ")
	if params[0] == paramAll && len(completedTasks) > 0 {
		result = append([]string{"Pending tasks:"}, result...)
		result = append(result, "Completed tasks:")
		for _, completedTask := range completedTasks {
			result = append(result, completedTask.String())
		}
	}

	return strings.Join(result, "\n")
}
