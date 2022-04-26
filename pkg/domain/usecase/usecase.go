package usecase

type UseCase interface {
	Execute(text string) string
}

type Factory interface {
	NewEchoUseCase() UseCase
}

type factory struct {
}

func NewFactory() Factory {
	return &factory{}
}

func (f *factory) NewEchoUseCase() UseCase {
	return &EchoUseCase{}
}

type EchoUseCase struct{}

func (e *EchoUseCase) Execute(text string) string {
	return text
}
