package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func newApiBot(token string) *tgbotapi.BotAPI {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}
