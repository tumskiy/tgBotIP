package thirdparty

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"tgBotIP/internal/database"
	"tgBotIP/internal/env"
	"tgBotIP/internal/types"
)

func Run() {
	parseEnvToken := env.ParseEnv("TELEGRAM_TOKEN")
	parseEnvErr := env.ParseEnv("ERROR_RESPONSE_MESSAGE")
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
				exist, err := database.ExistUser(update.Message.Chat.ID)
				if err != nil {
					tgbotapi.NewMessage(update.Message.Chat.ID, parseEnvErr)
				}
				if exist {
					_, err = bot.Send(sendIP(update.Message.Chat.ID, update.Message.MessageID))
					if err != nil {
						return
					}
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					msg.ReplyToMessageID = update.Message.MessageID
					_, err := bot.Send(sendError(update.Message.Chat.ID, update.Message.MessageID))
					if err != nil {
						return
					}
					continue
				}

			case "password":
				parseEnvParse := env.ParseEnv("ACCESS_PASSWORD")
				password := strings.TrimPrefix(update.Message.Text, "/password ")
				user := types.User{
					ID:       update.Message.From.ID,
					Name:     update.Message.From.UserName,
					Replay:   update.Message.MessageID,
					Password: password,
					ChatID:   update.Message.Chat.ID,
				}
				if password == parseEnvParse {
					err := database.CreateUser(user)
					msg := tgbotapi.NewMessage(user.ChatID, update.Message.Text)
					msg.ReplyToMessageID = user.Replay
					_, err = bot.Send(sendSuccess(update.Message.Chat.ID, update.Message.MessageID))
					if err != nil {
						return
					}
					if err != nil {
						return
					}
					continue
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					msg.ReplyToMessageID = update.Message.MessageID
					_, err := bot.Send(sendError(update.Message.Chat.ID, update.Message.MessageID))
					if err != nil {
						return
					}
				}
				continue

			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID
				_, err := bot.Send(sendError(update.Message.Chat.ID, update.Message.MessageID))
				if err != nil {
					return
				}
				continue
			}
		}
	}
}

func sendError(chatID int64, replay int) tgbotapi.MessageConfig {
	parseEnvErr := env.ParseEnv("ERROR_RESPONSE_MESSAGE")
	messageConfig := tgbotapi.NewMessage(chatID, parseEnvErr)
	messageConfig.ReplyToMessageID = replay
	return messageConfig
}

func sendIP(chatID int64, replay int) tgbotapi.MessageConfig {
	info, err := Request()
	if err != nil {
		return tgbotapi.NewMessage(chatID, "Failed to retrieve IP information")
	}
	messageText := fmt.Sprintf(`
ip: %s
city: %s
provider: %s
`, info.IP, info.City, info.AsnOrg)
	msg := tgbotapi.NewMessage(chatID, messageText)
	msg.ReplyToMessageID = replay
	return msg
}

func sendSuccess(chatID int64, replay int) tgbotapi.MessageConfig {
	messageConfig := tgbotapi.NewMessage(chatID, "User create")
	messageConfig.ReplyToMessageID = replay
	return messageConfig
}
