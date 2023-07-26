package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5327315256:AAGvkuzryn2dFx0NmYgRpajAWpjAjS_HoNs")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.IsCommand() {
			var msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = "Binance"
			bot.Send(msg);
		}
	}
}
