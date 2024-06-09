package interaction

import (
	"fmt"
	"strings"
	"tgBotIP/internal/database"
	"tgBotIP/internal/env"
	"tgBotIP/internal/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler interface {
	HandleIP(*tgbotapi.BotAPI, tgbotapi.Update)()
	HandlePassword(*tgbotapi.BotAPI, tgbotapi.Update)()
	HandleDefault(*tgbotapi.BotAPI, tgbotapi.Update)()

	Request() (types.Site, error)
}

var d database.Databaser

func (r *Req) HandleIP(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	parseEnvErr := env.ParseEnv("ERROR_RESPONSE_MESSAGE")
	fmt.Println("select exist user")
	exist, err := d.ExistUser(update.Message.From.ID)
	if err != nil {
		tgbotapi.NewMessage(update.Message.Chat.ID, parseEnvErr)
	}
	if exist {
		fmt.Println("user exist")
		_, err = bot.Send(r.sendIP(update.Message.Chat.ID, update.Message.MessageID))
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

func (r *Req) HandlePassword(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
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
		err := d.CreateUser(user)
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

func (r *Req) HandleDefault(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
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

func (r *Req) sendIP(chatID int64, replay int) tgbotapi.MessageConfig {
	info, err := r.Request()
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
