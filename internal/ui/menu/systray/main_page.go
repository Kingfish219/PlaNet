package systray

import (
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

	dnsConfigPage := NewDnsConfigPage(mainPage.systray)

	itemList := []ui.Item{
		{
			Key:   "systray_main_dns_config",
			Title: "Config",
			Page:  dnsConfigPage.Initialize(),
		},
		{
			Key:   "systray_main_dns_set",
			Title: "Set",
			Exec: func() any {
				setConfig(mainPage.systray)
				return nil
			},
		},
		{
			Key:   "systray_main_dns_reset",
			Title: "Reset",
			Exec: func() any {
				resetConfig(mainPage.systray)
				return nil
			},
		},
		{
			Key:   "systray_main_dns_delete",
			Title: "Delete This Config",
			Exec: func() any {
				deleteConfig(mainPage.systray)
				return nil
			},
		},
	}

	return &ui.Page{
		Key:   "systray_main_dns",
		Items: itemList,
	}
}
