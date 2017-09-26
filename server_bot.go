package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"telehorn/config"
	logrus "github.com/sirupsen/logrus"
)

func init(){
	config.InitLog("bot")
}

func StartBot() {
	bot, err := tgbotapi.NewBotAPI(config.BotKey())
	if err != nil {
		logrus.Panic(err)
	}

	bot.Debug = true

	logrus.Info("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Chat.Type == "group" {
			// TODO: Group chat response

		} else {
			// TODO: Personal chat response
		}

		logrus.Info("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
