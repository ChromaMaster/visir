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

	When("the echo command is received", func() {
		It("executes the echo command", func() {
			tgbotAPI := &TelegramBotSeam{
				BotAPI:  nil,
				updates: []tgbotapi.Update{echoUpdate()},
			}

			telegramBot := telegrambot.NewBot(tgbotAPI, &usecaseFactory)

			telegramBot.Start()

			Expect(usecaseFactory.ExecCount).To(Equal(1))
			Expect(len(tgbotAPI.messagesSent)).To(Equal(1))
		})
	})

	When("the public_ip command is received", func() {
		It("executes the public_ip command", func() {
			tgbotAPI := &TelegramBotSeam{
				BotAPI:  nil,
				updates: []tgbotapi.Update{publicIp()},
			}

			telegramBot := telegrambot.NewBot(tgbotAPI, &usecaseFactory)

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

func publicIp() tgbotapi.Update {
	return tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "/publicIp",
			Chat: &tgbotapi.Chat{
				ID: int64(0),
			},
		},
	}
}
