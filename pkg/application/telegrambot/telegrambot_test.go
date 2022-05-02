package telegrambot_test

import (
	"github.com/ChromaMaster/visir/pkg/application/telegrambot"
	"github.com/ChromaMaster/visir/pkg/domain/usecase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type FakeUsecaseFactory struct {
	ExecCount int
}

func (f *FakeUsecaseFactory) NewEchoUseCase() usecase.UseCase {
	return &FakeUsecase{factory: f}
}

func (f *FakeUsecaseFactory) NewPublicIpUseCase() usecase.UseCase {
	return &FakeUsecase{factory: f}
}

type FakeUsecase struct {
	factory *FakeUsecaseFactory
}

func (f FakeUsecase) Execute(text string) string {
	f.factory.ExecCount++
	return ""
}

var _ = Describe("Telegrambot", func() {

	usecaseFactory := FakeUsecaseFactory{}

	BeforeEach(func() {
		usecaseFactory.ExecCount = 0
	})

	When("a message from a unauthorized user is received", func() {
		It("it's not handled", func() {
			tgbotAPI := &TelegramBotSeam{
				BotAPI:  nil,
				updates: []tgbotapi.Update{unauthorizedMessage()},
			}

			telegramBot := telegrambot.NewBot(tgbotAPI, 0, &usecaseFactory)

			telegramBot.Start()

			Expect(usecaseFactory.ExecCount).To(Equal(0))
			Expect(len(tgbotAPI.messagesSent)).To(Equal(0))
		})
	})

	When("the echo command is received", func() {
		It("executes the echo command", func() {
			tgbotAPI := &TelegramBotSeam{
				BotAPI:  nil,
				updates: []tgbotapi.Update{echoUpdate()},
			}

			telegramBot := telegrambot.NewBot(tgbotAPI, 0, &usecaseFactory)

			telegramBot.Start()

			Expect(usecaseFactory.ExecCount).To(Equal(1))
			Expect(len(tgbotAPI.messagesSent)).To(Equal(1))
		})
	})

	When("the public_ip command is received", func() {
		It("executes the public_ip command", func() {
			tgbotAPI := &TelegramBotSeam{
				BotAPI:  nil,
				updates: []tgbotapi.Update{publicIpUpdate()},
			}

			telegramBot := telegrambot.NewBot(tgbotAPI, 0, &usecaseFactory)

			telegramBot.Start()

			Expect(usecaseFactory.ExecCount).To(Equal(1))
			Expect(len(tgbotAPI.messagesSent)).To(Equal(1))
		})
	})
})

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

func unauthorizedMessage() tgbotapi.Update {
	msg := tgMessage("/echo foo")
	msg.Message.From.ID = 1
	return msg
}

func echoUpdate() tgbotapi.Update {
	return tgMessage("/echo foo")
}

func publicIpUpdate() tgbotapi.Update {
	return tgMessage("/publicIp")
}

func tgMessage(text string) tgbotapi.Update {
	return tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: text,
			Chat: &tgbotapi.Chat{
				ID: int64(0),
			},
			From: &tgbotapi.User{
				ID: 0,
			},
		},
	}
}
