package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tgb/internal/stores"
)

type TgBot struct {
	api *tgbotapi.BotAPI
	db  *stores.Store
}

func New(token string) *TgBot {
	db, err := stores.New()
	if err != nil {
		log.Fatal(err)
	}

	return &TgBot{
		api: newApiBot(token),
		db:  db,
	}
}

func (r *TgBot) Run() {
	r.api.Debug = true
	log.Printf("Authorized on account %s", r.api.Self.UserName)
}
