package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageSender(message string, id int64) {
	Bot.Send(
		tgbotapi.NewMessage(
			id, message,
		),
	)
	log.Println("Message sent successfully")
}
