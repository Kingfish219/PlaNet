package io

import "fmt"

type ConsoleIO struct {
}

func (console ConsoleIO) Initialize() error {
	_, err := fmt.Println("Welcome")

	return err
}
