package handlers

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"tasks/internal/storage"
)

func updateFunc(args string) string {
	params := strings.Split(args, " ")
	if len(params) == 0 {
		return errors.New("Invalid number of arguments are given").Error()
	}
	id, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil {
		return err.Error()
	}
	err = storage.Update(uint(id), strings.Join(params[1:], " "))
	if err != nil {
		return err.Error()
	}
	task, err := storage.Get(uint(id))
	return fmt.Sprintf("Task %v is updated", task)
}
