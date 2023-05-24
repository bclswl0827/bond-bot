package main

import (
	"log"
	"time"

	workingday "github.com/Admingyu/go-workingday"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	var (
		args Args
		conf Config
	)

	args.ReadFlags()
	conf.ReadConfig(args.Path)

	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Authorized account", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	var flag bool
	go ResetFlag(conf, &flag)

	for {
		work, _ := workingday.IsWorkDay(
			time.Now(), "CN",
		)
		if work && !flag {
			MessageSender(
				bot, conf.ChatId, BondFilter(
					BondParser(BondData()),
				),
			)

			flag = true
		}

		time.Sleep(time.Second)
	}
}
