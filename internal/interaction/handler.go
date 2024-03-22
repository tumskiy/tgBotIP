package interaction

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"tgBotIP/internal/database"
	"tgBotIP/internal/env"
	"tgBotIP/internal/types"
)

func handleIP(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	parseEnvErr := env.ParseEnv("ERROR_RESPONSE_MESSAGE")
	fmt.Println("select exist user")
	exist, err := database.ExistUser(update.Message.From.ID)
	if err != nil {
		tgbotapi.NewMessage(update.Message.Chat.ID, parseEnvErr)
	}
	if exist {
		fmt.Println("user exist")
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
	}

}

func handlePassword(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	parseEnvPassword := env.ParseEnv("ACCESS_PASSWORD")
	password := strings.TrimPrefix(update.Message.Text, "/password ")
	user := types.User{
		ID:       update.Message.From.ID,
		Name:     update.Message.From.UserName,
		Replay:   update.Message.MessageID,
		Password: password,
		ChatID:   update.Message.Chat.ID,
	}
	if password == parseEnvPassword {
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
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		_, err := bot.Send(sendError(update.Message.Chat.ID, update.Message.MessageID))
		if err != nil {
			return
		}
	}
}

func handleDefault(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID
	_, err := bot.Send(sendError(update.Message.Chat.ID, update.Message.MessageID))
	if err != nil {
		return
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
