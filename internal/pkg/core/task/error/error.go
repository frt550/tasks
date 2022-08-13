package error

import (
	"log"
	"tasks/internal/config"

	"github.com/pkg/errors"
)

const messageInternalError = "Internal error"

var TaskError = errors.New("")

func Error(err error) string {
	if errors.Is(err, TaskError) {
		return err.Error()
	} else {
		if config.Config.App.Debug == "true" {
			return err.Error()
		} else {
			log.Println(err)
			return messageInternalError
		}
	}
}
