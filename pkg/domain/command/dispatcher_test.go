package command_test

import (
	"github.com/ChromaMaster/visir/pkg/domain/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dispatcher", func() {
	var (
		repository *command.Repository
	)

	BeforeEach(func() {
		repository = command.NewRepository()
	})

	When("a command belongs to the dispatcher", func() {
		It("can be executed", func() {
			// TODO(fede): Esto tiene sentido aquí o mejor en el before each?
			err := repository.RegisterCommand(command.Command{Name: "command1", Command: func(string) error {
				return nil
			}})
			Expect(err).ToNot(HaveOccurred())

			dispatcher:= command.NewDispatcher(repository)
			err = dispatcher.Run("command1")
			Expect(err).ToNot(HaveOccurred())
		})
	})

	When("a command does not belong to the dispatcher", func() {
		It("cannot be executed", func() {
			dispatcher := command.Dispatcher{Commands: repository}
			err := dispatcher.Run("command2")
			Expect(err).To(MatchError(command.ErrCommandNotFound))
		})
	})
})
