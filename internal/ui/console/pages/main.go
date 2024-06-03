package pages

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/ui"
)

func Main(repo interfaces.DnsRepository) *ui.Page {
	page := &ui.Page{
		Key:   "c_main",
		Title: "What do you want to do?",
		Items: []ui.Item{
			{
				Key:      "c_main_dns",
				Title:    "1. DNS Management",
				ShortKey: "1",
				Exec: func() {
					fmt.Println("")
				},
			},
			{
				Key:      "c_main_ni",
				Title:    "2. Network Interface Management",
				ShortKey: "2",
			},
			{
				Key:      "c_main_tools",
				Title:    "3. Tools",
				ShortKey: "3",
			},
		},
	}

	page.Items[0].Page = DnsManagement(page, repo)

	return page
}
