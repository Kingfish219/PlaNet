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
		Items: []ui.Item{
			{
				Title: "DNS Management",
				Exec: func() {
					fmt.Println("Test")
				},
			},
		},
	}
}