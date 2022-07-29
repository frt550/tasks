package handlers

import (
	"strings"
	"tasks/internal/storage"
)

func showFunc(args string) string {
	data := storage.List()
	pending := make([]string, 0, len(data))
	completed := make([]string, 0, len(data))
	for _, v := range data {
		if !v.IsCompleted() {
			pending = append(pending, v.String())
		} else {
			completed = append(completed, v.String())
		}
	}

	var result = ""
	if len(pending) == 0 {
		result = "All tasks are done!\n"
	} else {
		result = strings.Join(pending, "\n")
	}

	params := strings.Split(args, " ")
	if len(params) == 1 && params[0] == "all" && len(completed) > 0 {
		result += "\n\nCompleted tasks:\n"
		result += strings.Join(completed, "\n")
	}
	return result
}
