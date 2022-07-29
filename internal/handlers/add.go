package handlers

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"tasks/internal/storage"
)

func addFunc(args string) string {
	args = strings.TrimSpace(args)
	if len(args) == 0 {
		return errors.New("Title should not be empty").Error()
	}
	task, err := storage.Add(args)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Task %v is added", task)
}
