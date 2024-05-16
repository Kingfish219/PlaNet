package console

import (
	"fmt"
	"os"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/Kingfish219/PlaNet/internal/ui"
)

type ConsoleUI struct {
	Name       string
	ActivePage ui.Page
}

func New() *ConsoleUI {
	return &ConsoleUI{}
}

func (console *ConsoleUI) Initialize() error {
	console.drawLogo()
	FeedMainPage(console)

	fmt.Println()
	_, err := fmt.Println("What do you want to do?")
	console.addActions()
	console.captureKeyboard()

	return err
}

func (console *ConsoleUI) captureKeyboard() {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.CtrlC:
			os.Exit(0)
			return true, nil
		case keys.RuneKey:
			fmt.Printf("\rYou pressed the rune key: %s\n", key)
		default:
			fmt.Printf("\rYou pressed: %s\n", key)
		}

		return false, nil
	})
}

func (console *ConsoleUI) drawLogo() {
	fmt.Println("=======================================")
	fmt.Println("           Welcome to PlaNet           ")
	fmt.Println("=======================================")
}

func (console *ConsoleUI) addActions() {
	for index, action := range console.ActivePage.Items {
		fmt.Println(fmt.Sprint(index+1)+".", action.Title)
	}
}
