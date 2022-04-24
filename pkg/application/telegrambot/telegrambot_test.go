package telegrambot_test

import (
	"github.com/ChromaMaster/visir/pkg/application/telegrambot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type FakeEchoUserCase struct {
	ExecCount int
}

func (f *FakeEchoUserCase) Execute(text string) string {
	f.ExecCount++
	return ""
}

type TelegramBotSeam struct {
	*tgbotapi.BotAPI
	updates      []tgbotapi.Update
	messagesSent []tgbotapi.Chattable
}

func (t *TelegramBotSeam) Send(chattable tgbotapi.Chattable) (tgbotapi.Message, error) {
	t.messagesSent = append(t.messagesSent, chattable)
	return tgbotapi.Message{}, nil
}

func (t *TelegramBotSeam) GetUpdatesChan(config tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error) {
	result := make(chan tgbotapi.Update)
	go func() {
		defer close(result)
		for _, update := range t.updates {
			result <- update
		}
	}()
	return result, nil
}

var _ = Describe("Telegrambot", func() {
	When("the echo command is received", func() {
		It("executes the echo command", func() {
			echoUsecase := &FakeEchoUserCase{}
			tgbotAPI := &TelegramBotSeam{
				BotAPI:  nil,
				updates: []tgbotapi.Update{echoUpdate()},
			}

			telegramBot := telegrambot.NewBot(tgbotAPI, echoUsecase)

			telegramBot.Start()

			Expect(echoUsecase.ExecCount).To(Equal(1))
			Expect(len(tgbotAPI.messagesSent)).To(Equal(1))
		})
	})
})

func echoUpdate() tgbotapi.Update {
	return tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "/echo foo",
			Chat: &tgbotapi.Chat{
				ID: int64(0),
			},
		},
	}
}
