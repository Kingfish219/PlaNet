package console

import (
	"fmt"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/Kingfish219/PlaNet/internal/ui"
	"github.com/inancgumus/screen"
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
	FeedUI(console)

	fmt.Println()
	fmt.Println("What do you want to do?")
	console.buildActivePage()

	return nil
}

func (console *ConsoleUI) buildActivePage() {
	console.clearConsole()
	console.drawLogo()
	console.buildUI(console.ActivePage)
	console.buildKeyboard(console.ActivePage)
}

func (console *ConsoleUI) clearConsole() {
	screen.Clear()
	screen.MoveTopLeft()
}

func (console *ConsoleUI) drawLogo() {
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println("                PlaNet                 ")
	fmt.Println("=======================================")
	fmt.Println()
}

func (console *ConsoleUI) buildUI(page ui.Page) {
	fmt.Println(page.Title)
	fmt.Println()
	for _, action := range page.Items {
		fmt.Println(action.Title)
	}
}

func (console *ConsoleUI) buildKeyboard(page ui.Page) {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		for _, item := range page.Items {
			if item.ShortKey == key.String() {
				if item.Page.Key == "" {
					console.ActivePage = item.Page
					console.buildActivePage()
				} else {
					item.Exec()
				}
			}
		}

		return true, nil
	})
}
