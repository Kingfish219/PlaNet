package console

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/ui"
)

type ConsoleUI struct {
	Name          string
	ActivePage    *ui.Page
	dnsRepository interfaces.DnsRepository
}

func New(dnsRepository interfaces.DnsRepository) *ConsoleUI {
	return &ConsoleUI{
		dnsRepository: dnsRepository,
	}
}

func (console *ConsoleUI) Initialize() error {
	console.drawLogo()
	page := FeedUI(console)
	fmt.Println()
	fmt.Println("What do you want to do?")
	console.BuildPage(page)

	return nil
}

func (console *ConsoleUI) BuildPage(page *ui.Page) {
	// console.clearConsole()
	console.drawLogo()
	console.buildUI(page)
	console.buildKeyboard(page)
	page.Initiate()
}

func (console *ConsoleUI) clearConsole() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (console *ConsoleUI) drawLogo() {
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println("                PlaNet                 ")
	fmt.Println("=======================================")
	fmt.Println()
}

func (console *ConsoleUI) buildUI(page *ui.Page) {
	fmt.Println(page.Title)
	fmt.Println()
	for _, action := range page.Items {
		fmt.Println(action.Title)
	}
}

func (console *ConsoleUI) buildKeyboard(page *ui.Page) {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		for _, item := range page.Items {
			if item.ShortKey == key.String() {
				if item.Page.Key != "" {
					console.BuildPage(item.Page)
				} else {
					item.Exec()
				}
			}
		}

		return true, nil
	})
}
