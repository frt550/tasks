package handlers

import (
	"fmt"
	"strconv"
	"tasks/internal/storage"
)

func completeFunc(args string) string {
	id, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		return err.Error()
	}
	err = storage.Complete(uint(id))
	if err != nil {
		return err.Error()
	}
	task, err := storage.Get(uint(id))
	return fmt.Sprintf("Task %v is completed", task)
}
