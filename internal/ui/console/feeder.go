package console

import (
	"fmt"
	"os"

	"github.com/Kingfish219/PlaNet/internal/ui"
)

func FeedUI(console *ConsoleUI) {
	console.ActivePage = MainPage()
}

func MainPage() ui.Page {
	return ui.Page{
		Key:   "main",
		Title: "What do you want to do?",
		Items: []ui.Item{
			{
				Key:      "main_dns",
				Title:    "1. DNS Management",
				ShortKey: "1",
				Page: ui.Page{
					Key:   "main_dns_dns",
					Title: "Dns Management",
					Items: []ui.Item{
						{
							Key:   "main_dns_dns_config",
							Title: "Select",
						},
					},
				},
			},
			{
				Key:      "main_ni",
				Title:    "2. Network Interface Management",
				ShortKey: "2",
			},
			{
				Key:      "main_tools",
				Title:    "3. Tools",
				ShortKey: "3",
			},
			{
				Key:      "main_exit",
				Title:    "0. Exit",
				ShortKey: "0",
				Exec: func() {
					fmt.Println()
					fmt.Println("Goodbye...")
					os.Exit(0)
				},
			},
		},
	}
}
