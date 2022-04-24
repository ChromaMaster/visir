package usecase_test

import (
	"github.com/ChromaMaster/visir/pkg/domain/usecase"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Echo Use Case", func() {
	It("returns the same text that was provided", func() {
		echoUsecase := usecase.NewEchoUseCase()
		result := echoUsecase.Execute("text")
		Expect(result).To(Equal("text"))
	})
})
