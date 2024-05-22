package systray

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/ui"
	"github.com/Kingfish219/PlaNet/network/dns"
	"github.com/Kingfish219/PlaNet/network/ni"
)

type DnsPage struct {
	systray        *SystrayUI
	key            string
	SetConfigTitle func(string)
}

func NewDnsPage(systray *SystrayUI, key string) *DnsPage {
	return &DnsPage{
		systray: systray,
		key:     key,
	}
}

func (dnsConfig *DnsPage) Initialize() ui.Page {

	dnsConfigPage := NewDnsConfigPage(dnsConfig.systray, dnsConfig.key+"_config")

	itemList := []ui.Item{
		{
			Key:   dnsConfig.key + "_config",
			Title: "Config",
			Page:  dnsConfigPage.Initialize(),
		},
		{
			Key:   dnsConfig.key + "_set",
			Title: "Set",
			Exec: func() {
				fmt.Print("Set")
				setConfig(dnsConfig.systray)
			},
		},
		{
			Key:   dnsConfig.key + "_reset",
			Title: "Reset",
			Exec: func() {
				fmt.Print("Reset")
				resetConfig(dnsConfig.systray)
			},
		},
		{
			Key:   dnsConfig.key + "_delete",
			Title: "Delete This Config",
			Exec: func() {
				fmt.Print("Delte")
				deleteConfig(dnsConfig.systray)

			},
		},
	}

	return ui.Page{
		Key:   dnsConfig.key,
		Items: itemList,
	}
}

func deleteConfig(systrayUI *SystrayUI) {
	activeInterfaceNames, err := ni.GetActiveNetworkInterface()
	if activeInterfaceNames == nil || err != nil {
		fmt.Println(err)

		return
	}

	if len(activeInterfaceNames) == 0 {
		fmt.Println("no active network interface found")

		return
	}

	var targetDns = systrayUI.selectedDnsConfiguration
	err = systrayUI.dnsRepository.DeleteDnsConfigurations(systrayUI.selectedDnsConfiguration)
	if err != nil {
		fmt.Println(err)
		return
	}

	for key, item := range systrayUI.DnsMenus {
		if key == targetDns.Name {
			item.Hide()
			//dnsConfigMenu.SetTitle("DNS config: -")
		}
	}
	delete(systrayUI.DnsMenus, targetDns.Name)
	dnsService := dns.DnsService{}
	_, err = dnsService.ChangeDns(dns.ResetDns, systrayUI.connectedDnsConfiguration)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(targetDns.Name + " deleted successfully.")

	systrayUI.setIcon(false)
	systrayUI.setToolTip("Not connected")
}

func resetConfig(systrayUI *SystrayUI) {
	fmt.Println(systrayUI.selectedDnsConfiguration)
	fmt.Println("dnsMenu", systrayUI.DnsMenus)
	dnsService := dns.DnsService{}
	_, err := dnsService.ChangeDns(dns.ResetDns, systrayUI.connectedDnsConfiguration)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Shecan disconnected successfully.")

	systrayUI.setIcon(false)
	systrayUI.setToolTip("Not connected")

}

func setConfig(systrayUI *SystrayUI) {
	dnsService := dns.DnsService{}
	_, err := dnsService.ChangeDns(dns.SetDns, systrayUI.selectedDnsConfiguration)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Shecan set successfully.")

	systrayUI.setIcon(true)
	systrayUI.setToolTip("Connected to: Shecan")
}
