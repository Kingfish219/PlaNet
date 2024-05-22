package systray

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/presets"
	"github.com/Kingfish219/PlaNet/internal/ui"
)

func Feed(systray *SystrayUI) error {
	err := setInitConfig(systray)
	if err != nil {
		return err
	}

	systray.Page = pages(systray)
	return nil
}

func pages(systray *SystrayUI) ui.Page {

	dnsPage := NewDnsPage(systray, "systray_main_dns")

	return ui.Page{
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
				Title: "Console",
				Exec: func() {
					fmt.Println("Console")
				},
			},
		},
	}
}

func setInitConfig(systrayUI *SystrayUI) error {
	dnsConfigurations, err := systrayUI.dnsRepository.GetDnsConfigurations()
	if err != nil {
		return err
	}

	if len(dnsConfigurations) == 0 {
		presetDnsList := presets.GetDnsPresets()
		for _, pre := range presetDnsList {
			systrayUI.dnsRepository.ModifyDnsConfigurations(pre)
		}

		dnsConfigurations, err = systrayUI.dnsRepository.GetDnsConfigurations()
		if err != nil {
			return err
		}
	}

	systrayUI.dnsConfigurations = dnsConfigurations
	return nil
}
