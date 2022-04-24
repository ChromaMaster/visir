package telegrambot_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTelegrambot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Telegrambot Suite")
}
