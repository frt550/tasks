package handlers

func helpFunc(args string) string {
	return `/help - list commands
/show - show pending tasks
/show all - show all tasks including completed
/add <title> - add new task
/complete <id> - complete task
/update <id> <title> - update title of task
/delete <id> - delete task
/read <id> - read detailed info about task`
}
