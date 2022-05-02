package main

import (
	"github.com/ChromaMaster/visir/pkg/application/telegrambot"
	"github.com/ChromaMaster/visir/pkg/domain/usecase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strconv"
)

func main() {
	adminID, err := strconv.Atoi(os.Getenv("VISIR_ADMIN_ID"))
	if err != nil {
		panic(err)
	}
	botToken := os.Getenv("VISIR_TGBOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}
	telegramBot := telegrambot.NewBot(bot, adminID, usecase.NewFactory())

	telegramBot.Start()
}
