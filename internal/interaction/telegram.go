package interaction

import (
	"log"
	"tgBotIP/internal/env"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Runner interface{
	Run()
}

func Run() {
	parseEnvToken := env.ParseEnv("TELEGRAM_TOKEN")
	bot, err := tgbotapi.NewBotAPI(parseEnvToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	var h Handler
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			switch update.Message.Command() {
			case "ip":
				h.HandleIP(bot, update)
				continue
			case "password":
				h.HandlePassword(bot, update)
				continue
			default:
				h.HandleDefault(bot, update)
				continue
			}
		}
	}
}
