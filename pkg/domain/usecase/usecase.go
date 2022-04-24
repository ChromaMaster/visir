package usecase

type EchoUseCase struct {
}

func (e *EchoUseCase) Execute(text string) string {
	return text
}

func NewEchoUseCase() *EchoUseCase {
	return &EchoUseCase{}
}
