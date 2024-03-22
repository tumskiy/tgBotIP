package main

import (
	"tgBotIP/internal/database"
	"tgBotIP/internal/env"
	"tgBotIP/internal/interaction"
)

func main() {
	database.Init()
	env.Init()
	interaction.Run()
}
