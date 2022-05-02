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
	adminID  int
	factory  usecase.Factory
}

func NewBot(tgbotAPI TelegramBot, adminID int, factory usecase.Factory) *Bot {
	return &Bot{
		tgbotAPI: tgbotAPI,
		adminID:  adminID,
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

		if update.Message.From.ID != b.adminID {
			continue
		}

		messageText := update.Message.Text
		fmt.Println(messageText)
		if !isCommand(messageText) {
			continue
		}
		comm, args := extractCommandInfo(messageText)
		var response string
		switch comm {
		case "echo":
			response = b.factory.NewEchoUseCase().Execute(args)
		case "publicIp":
			response = b.factory.NewPublicIpUseCase().Execute("")
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ReplyToMessageID = update.Message.MessageID
		_, err := b.tgbotAPI.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func isCommand(text string) bool {
	if !strings.HasPrefix(text, "/") {
		return false
	}
	if len(strings.ReplaceAll(text, "/", "")) == 0 {
		return false
	}
	return true
}

func extractCommandInfo(text string) (command string, args string) {
	fields := strings.Fields(strings.ReplaceAll(text, "/", ""))
	command = fields[0]
	args = strings.Join(fields[1:], "")
	return
}
