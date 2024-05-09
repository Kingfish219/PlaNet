package console

import "fmt"

func FeedActions(console *ConsoleUI) {
	console.Actions = append(console.Actions, Dns())
}

func Dns() Action {
	return Action{
		Title: "DNS Management",
		Exec: func() {
			fmt.Println("Test")
		},
	}
}
