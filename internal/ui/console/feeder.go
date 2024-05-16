package console

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/ui"
)

func FeedMainPage(console *ConsoleUI) {
	console.ActivePage = MainPage()
}

func MainPage() ui.Page {
	return ui.Page{
		Key: "main",
		Items: []ui.Item{
			{
				Key:   "main_dns",
				Title: "DNS Management",
				Page: ui.Page{
					Key: "main_dns_dns",
				},
			},
			{
				Title: "Network Interface Management",
				Exec: func() {
					fmt.Println("Test")
				},
			},
			{
				Title: "Tools",
				Exec: func() {
					fmt.Println("Test")
				},
			},
			{
				Title: "Exit",
				Exec: func() {
					fmt.Println("Test")
				},
			},
		},
	}
}
