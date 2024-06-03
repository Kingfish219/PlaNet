package systray

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/publisher"
	"github.com/Kingfish219/PlaNet/internal/ui"
	"github.com/Kingfish219/PlaNet/internal/ui/console"
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
				Exec: func() {
					fmt.Println("Network Interface Management")
				},
			},
			{
				Title: "Tools",
				Exec: func() {
					fmt.Println("Tools")
				},
			},
			{
				Title: "Console",
				Exec: func() {
					console := console.New(mainPage.systray.dnsRepository)
					publisher := publisher.Publisher{}
					publisher.UISubscribers = append(publisher.UISubscribers, console)
					publisher.PublishUI("main")
				},
			},
		},
	}
}
