package main

import (
	"flag"
	"log"
	"os"
	"time"

	workingday "github.com/Admingyu/go-workingday"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// 先将就用吧
var (
	// 配置文件路径
	conf string
	// 配置文件内容
	config Config
	// Bot 实例
	Bot *tgbotapi.BotAPI
	// 记录推送
	Sent bool
)

func main() {
	// 命令行参数
	flag.StringVar(&conf, "c", "./config.json", "Config file path")
	flag.Parse()

	// 读取配置文件
	ReadConfig(conf)

	// 设定代理
	if len(config.Proxy) > 0 {
		// 设定环境变量
		os.Setenv("HTTPS_PROXY", config.Proxy)
	}

	// 初始化 Bot
	if bot, err := tgbotapi.NewBotAPI(config.Token); err != nil {
		log.Fatalln(err)
	} else {
		Bot = bot
		log.Println("Authorized on account", Bot.Self.UserName)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	// 定时重置 Flag
	go ResetFlag()

	// 定时发送讯息
	for {
		if work, _ := workingday.IsWorkDay(
			time.Now(), "CN",
		); work && !Sent {
			MessageSender(
				BondFilter(
					BondParser(
						BondData(),
					),
				), config.ChatId,
			)
			Sent = true
		}
	}
}
