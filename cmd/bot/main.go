package main

import (
	"log"
	botPkg "tasks/internal/pkg/bot"
	cmdAddPkg "tasks/internal/pkg/bot/command/add"
	cmdCompletePkg "tasks/internal/pkg/bot/command/complete"
	cmdDeletePkg "tasks/internal/pkg/bot/command/delete"
	cmdHelpPkg "tasks/internal/pkg/bot/command/help"
	cmdReadPkg "tasks/internal/pkg/bot/command/read"
	cmdShowPkg "tasks/internal/pkg/bot/command/show"
	cmdUpdatePkg "tasks/internal/pkg/bot/command/update"
	poolPkg "tasks/internal/pkg/core/pool"
	taskPkg "tasks/internal/pkg/core/task"
	"tasks/internal/pkg/core/task/repository/postgres"
)

func main() {
	var task = taskPkg.New(postgres.New(poolPkg.GetInstance()))
	runBot(task)
}

func runBot(task taskPkg.Interface) {
	var bot botPkg.Interface
	{
		bot = botPkg.MustNew()

		commandAdd := cmdAddPkg.New(task)
		bot.RegisterHandler(commandAdd)

		commandShow := cmdShowPkg.New(task)
		bot.RegisterHandler(commandShow)

		commandComplete := cmdCompletePkg.New(task)
		bot.RegisterHandler(commandComplete)

		commandUpdate := cmdUpdatePkg.New(task)
		bot.RegisterHandler(commandUpdate)

		commandDelete := cmdDeletePkg.New(task)
		bot.RegisterHandler(commandDelete)

		commandRead := cmdReadPkg.New(task)
		bot.RegisterHandler(commandRead)

		commandHelp := cmdHelpPkg.New(map[string]string{
			commandAdd.Name():      commandAdd.Description(),
			commandShow.Name():     commandShow.Description(),
			commandComplete.Name(): commandComplete.Description(),
			commandUpdate.Name():   commandUpdate.Description(),
			commandDelete.Name():   commandDelete.Description(),
			commandRead.Name():     commandRead.Description(),
		})
		bot.RegisterHandler(commandHelp)
	}

	if err := bot.Run(); err != nil {
		log.Panic(err)
	}
}
