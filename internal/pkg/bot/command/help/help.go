package help

import (
	"fmt"
	"strings"
	commandPkg "tasks/internal/pkg/bot/command"
)

func New(extendedMap map[string]string) commandPkg.Interface {
	if extendedMap == nil {
		extendedMap = map[string]string{}
	}
	return &command{
		extended: extendedMap,
	}
}

type command struct {
	extended map[string]string
}

func (c *command) Name() string {
	return "help"
}

func (c *command) Description() string {
	return "prints list of all commands"
}

func (c *command) Process(_ string) string {
	result := []string{
		fmt.Sprintf("/%s - %s", c.Name(), c.Description()),
	}
	for cmd, description := range c.extended {
		result = append(result, fmt.Sprintf("/%s - %s", cmd, description))
	}
	return strings.Join(result, "\n")
}
