package handlers

import (
	"tasks/internal/commander"
)

const (
	helpCmd     = "help"
	showCmd     = "show"
	addCmd      = "add"
	completeCmd = "complete"
	deleteCmd   = "delete"
	updateCmd   = "update"
	readCmd     = "read"
)

func AddHandlers(c *commander.Commander) {
	c.RegisterHandler(helpCmd, helpFunc)
	c.RegisterHandler(showCmd, showFunc)
	c.RegisterHandler(addCmd, addFunc)
	c.RegisterHandler(completeCmd, completeFunc)
	c.RegisterHandler(deleteCmd, deleteFunc)
	c.RegisterHandler(updateCmd, updateFunc)
	c.RegisterHandler(readCmd, readFunc)
}
