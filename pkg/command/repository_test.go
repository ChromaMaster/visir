package command_test

import (
	"github.com/ChromaMaster/visir/pkg/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository", func() {
	var (
		repository *command.CommandRepository
	)

	BeforeEach(func() {
		repository = command.NewRepository()
	})

	It("is able to register a command", func() {
		err := repository.Register(command.Command{Name: "", Command: func(string) error { return nil }})

		Expect(err).ToNot(HaveOccurred())
	})

	It("is able to retrieve a command", func() {
		err := repository.Register(command.Command{Name: "command1", Command: func(string) error { return nil }})
		Expect(err).ToNot(HaveOccurred())

		command, err := repository.Get("command1")
		Expect(err).ToNot(HaveOccurred())
		Expect(command.Name).To(Equal("command1"))
	})

	When("a command with the same name is already registered", func() {
		It("cannot be registered", func() {
			err := repository.Register(command.Command{Name: "command1", Command: func(string) error { return nil }})
			Expect(err).ToNot(HaveOccurred())

			err = repository.Register(command.Command{Name: "command1", Command: func(string) error { return nil }})
			Expect(err).To(MatchError(command.ErrCommandAlreadyRegistered))
		})
	})
})
