package systray

import (
	"github.com/Kingfish219/PlaNet/internal/presets"
)

func Feed(systray *SystrayUI) error {
	err := setInitConfig(systray)
	if err != nil {
		return err
	}

	mainPage := NewMainPage(systray)
	systray.Page = mainPage.Initialize()
	return nil
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
