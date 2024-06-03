package systray

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/ui"
)

type MainPage struct {
	systray *SystrayUI
}

func NewMainPage(systray *SystrayUI) *MainPage {
	return &MainPage{
		systray: systray,
	}
}

func (mainPage *MainPage) Initialize() *ui.Page {

	dnsPage := NewDnsPage(mainPage.systray)

	return &ui.Page{
		Key: "systray_main",
		Items: []ui.Item{
			{
				Key:   "systray_main_dns",
				Title: "DNS Management",
				Page:  dnsPage.Initialize(),
			},
			{
				Title: "Network Interface Management",
				Exec: func() any {
					fmt.Println("Network Interface Management")
					return nil
				},
			},
			{
				Title: "Tools",
				Exec: func() any {
					fmt.Println("Tools")
					return nil
				},
			},
			{
				Title: "Console",
				Exec: func() any {
					fmt.Println("Console")
					return nil
				},
			},
		},
	}
}
