package main

import (
	"os"
	"tgb/internal/bot"
)

func main() {

	token := os.Getenv("TELEGRAM_BOT_API_TOKEN")
	tBot := bot.New(token)
	tBot.Run()
}
