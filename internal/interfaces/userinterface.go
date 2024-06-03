package interfaces

type UserInterface interface {
	Initialize() error
	Consume(command string) error
}
