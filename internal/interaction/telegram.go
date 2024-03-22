package interaction

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tgBotIP/internal/env"
)

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

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			switch update.Message.Command() {
			case "ip":
				handleIP(bot, update)
				continue
			case "password":
				handlePassword(bot, update)
				continue
			default:
				handleDefault(bot, update)
				continue
			}
		}
	}
}
