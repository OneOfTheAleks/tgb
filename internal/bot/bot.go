package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type TgBot struct {
	api *tgbotapi.BotAPI
}

func New(token string) *TgBot {
	return &TgBot{
		api: newApiBot(token),
	}
}

func (r *TgBot) Run() {
	r.api.Debug = true
	log.Printf("Authorized on account %s", r.api.Self.UserName)
}
