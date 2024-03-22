package ui

import "fmt"

type ConsoleUI struct {
}

func (console ConsoleUI) Initialize() error {
	_, err := fmt.Println("Welcome")

	return err
}
