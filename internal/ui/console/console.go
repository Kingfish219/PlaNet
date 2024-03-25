package console

import "fmt"

type ConsoleUI struct {
}

func New() *ConsoleUI {
	return &ConsoleUI{}
}

func (console *ConsoleUI) Initialize() error {
	_, err := fmt.Println("Welcome")

	return err
}
