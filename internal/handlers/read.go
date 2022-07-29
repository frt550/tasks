package handlers

import (
	"fmt"
	"strconv"
	"tasks/internal/storage"
	"time"
)

func readFunc(args string) string {
	id, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		return err.Error()
	}
	task, err := storage.Get(uint(id))
	if err != nil {
		return err.Error()
	}
	isCompleted := ""
	if task.IsCompleted() {
		isCompleted = "It is completed at " + task.GetCompletedAt().Format(time.RFC850)
	} else {
		isCompleted = "It is not completed yet"
	}

	return fmt.Sprintf(
		"You are viewing detailed info of task\n%v.\nThis task was created at %s.\n%s",
		task.String(),
		task.GetCreatedAt().Format(time.RFC850),
		isCompleted,
	)
}
