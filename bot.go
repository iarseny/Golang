package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"fmt"

	"net/http"
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
			println(update.Message.Text);
		} else {
			if update.Message != nil {
				var resp, err = http.Get("http://example.com/")
				fmt.Println(resp, err);
				fmt.Print(update.Message.Text)
			}
		}
	}
}
