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
	"github.com/Kingfish219/PlaNet/internal/ui/console/pages"
)

type ConsoleUI struct {
	Name          string
	ActivePage    *ui.Page
	ExitFunc      func()
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
	console.BuildPage(page)

	return nil
}

func (console *ConsoleUI) Consume(command string) error {
	switch command {
	case "main":
		fmt.Println("main")
		break
	case "new-config":
		fmt.Println("Console new config")
		break
	}

	return nil
}

func (console *ConsoleUI) BuildPage(page *ui.Page) {
	console.ActivePage = page

	console.clearConsole()
	console.drawLogo()
	console.buildUI(page)
	if page.Initiate != nil {
		page.Initiate()
	}

	if page.Parent != nil {
		fmt.Println()
		fmt.Println("0. Back")
		console.ExitFunc = console.setBackKey
	} else {
		fmt.Println()
		fmt.Println("0. Exit")
		console.ExitFunc = console.setExitKey
	}

	console.buildKeyboard(page)
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

func (console *ConsoleUI) setBackKey() {
	console.BuildPage(console.ActivePage.Parent)
}

func (console *ConsoleUI) setExitKey() {
	console.BuildPage(pages.Exit())
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
	if page.Items == nil {
		return
	}

	for _, action := range page.Items {
		fmt.Println(action.Title)
	}
}

func (console *ConsoleUI) buildKeyboard(page *ui.Page) {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		for _, item := range page.Items {
			if item.ShortKey == key.String() {
				if item.Page != nil {
					if item.Page.Key != "" {
						console.BuildPage(item.Page)
					}
				} else {
					item.Exec()
				}
			}
		}

		if key.String() == "0" {
			console.ExitFunc()
		}

		return false, nil
	})
}
