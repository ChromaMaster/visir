package telegrambot

import (
	"fmt"
	"github.com/ChromaMaster/visir/pkg/domain/usecase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

type TelegramBot interface {
	Send(chattable tgbotapi.Chattable) (tgbotapi.Message, error)
	GetUpdatesChan(config tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error)
}

type Bot struct {
	tgbotAPI TelegramBot
	factory  usecase.Factory
}

func NewBot(tgbotAPI TelegramBot, factory usecase.Factory) *Bot {
	return &Bot{
		tgbotAPI: tgbotAPI,
		factory:  factory,
	}
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.tgbotAPI.GetUpdatesChan(u)
	if err != nil {
		log.Println(err)
		return
	}

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}

		messageText := update.Message.Text
		fmt.Println(messageText)
		if strings.HasPrefix(messageText, "/echo") {
			args := strings.ReplaceAll(messageText, "/echo ", "")
			text := b.factory.NewEchoUseCase().Execute(args)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyToMessageID = update.Message.MessageID
			_, err := b.tgbotAPI.Send(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
