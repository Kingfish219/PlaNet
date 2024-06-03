package pages

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/ui"
)

func Main(repo interfaces.DnsRepository) *ui.Page {
	return &ui.Page{
		Key:   "c_main",
		Title: "What do you want to do?",
		Items: []ui.Item{
			{
				Key:      "c_main_dns",
				Title:    "1. DNS Management",
				ShortKey: "1",
				Page:     DnsManagement(repo),
				Exec: func() any {
					fmt.Println("")
					return nil
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
			{
				Key:      "c_main_exit",
				Title:    "0. Exit",
				ShortKey: "0",
				Page:     Exit(),
			},
		},
	}
}
