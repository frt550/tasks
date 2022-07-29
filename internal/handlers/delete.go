package handlers

import (
	"fmt"
	"strconv"
	"tasks/internal/storage"
)

func deleteFunc(args string) string {
	id, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		return err.Error()
	}
	task, err := storage.Get(uint(id))
	if err != nil {
		return err.Error()
	}
	err = storage.Delete(uint(id))
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Task %v is deleted", task)
}
