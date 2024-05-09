package console

import (
	"fmt"
)

type ConsoleUI struct {
	Name    string
	Actions []Action
}

func New() *ConsoleUI {
	return &ConsoleUI{}
}

func (console *ConsoleUI) Initialize() error {
	console.drawLogo()
	FeedActions(console)

	fmt.Println()
	_, err := fmt.Println("What do you want to do?")
	console.addActions()

	return err
}

func (console *ConsoleUI) drawLogo() {
	fmt.Println("=======================================")
	fmt.Println("           Welcome to PlaNet           ")
	fmt.Println("=======================================")
}

func (console *ConsoleUI) addActions() {
	for index, action := range console.Actions {
		fmt.Println(index+1, action.Title)
	}
}
