package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageSender(bot *tgbotapi.BotAPI, id int64, message string) {
	bot.Send(
		tgbotapi.NewMessage(
			id, message,
		),
	)

	log.Println("Message sent successfully")
}
