package main

import (
	"tgBotIP/internal/database"
	"tgBotIP/internal/env"
	"tgBotIP/internal/thirdparty"
)

func main() {
	database.Init()
	env.Init()
	thirdparty.Run()
}
