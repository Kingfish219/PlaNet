package systray

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/ui"
	"github.com/Kingfish219/PlaNet/network/dns"
	"github.com/Kingfish219/PlaNet/network/ni"
)

type DnsPage struct {
	systray *SystrayUI
}

func NewDnsPage(systray *SystrayUI) *DnsPage {
	return &DnsPage{
		systray: systray,
	}
}

func (dnsConfig *DnsPage) Initialize() *ui.Page {

	dnsConfigPage := NewDnsConfigPage(dnsConfig.systray)

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
				setConfig(dnsConfig.systray)
				return nil
			},
		},
		{
			Key:   "systray_main_dns_reset",
			Title: "Reset",
			Exec: func() any {
				resetConfig(dnsConfig.systray)
				return nil
			},
		},
		{
			Key:   "systray_main_dns_delete",
			Title: "Delete This Config",
			Exec: func() any {
				deleteConfig(dnsConfig.systray)
				return nil
			},
		},
	}

	return &ui.Page{
		Key:   "systray_main_dns",
		Items: itemList,
	}
}

func deleteConfig(systrayUI *SystrayUI) {
	activeInterfaceNames, err := ni.GetActiveNetworkInterface()
	if activeInterfaceNames == nil || err != nil {
		fmt.Printf("Error: %v \n", err)

		return
	}

	if len(activeInterfaceNames) == 0 {
		fmt.Println("no active network interface found")

		return
	}

	var targetDns = systrayUI.selectedDnsConfiguration
	err = systrayUI.dnsRepository.DeleteDnsConfigurations(systrayUI.selectedDnsConfiguration)
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		return
	}

	for key, item := range systrayUI.SystrayMenuItem {
		if key == "systray_main_dns_config_"+targetDns.Name {
			item.Hide()
			//dnsConfigMenu.SetTitle("DNS config: -")
		}
	}
	delete(systrayUI.SystrayMenuItem, "systray_main_dns_config_"+targetDns.Name)
	dnsService := dns.DnsService{}
	_, err = dnsService.ChangeDns(dns.ResetDns, systrayUI.connectedDnsConfiguration)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(targetDns.Name + " deleted successfully.")
	systrayUI.SystrayMenuItem["systray_main_dns_config"].SetTitle("Config")
	systrayUI.setIcon(false)
	systrayUI.setToolTip("Not connected")
}

func resetConfig(systrayUI *SystrayUI) {
	dnsService := dns.DnsService{}
	conectedDnsName := systrayUI.connectedDnsConfiguration.Name
	_, err := dnsService.ChangeDns(dns.ResetDns, systrayUI.connectedDnsConfiguration)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(conectedDnsName + " disconnected successfully.")

	systrayUI.setIcon(false)
	systrayUI.setToolTip("Not connected")

}

func setConfig(systrayUI *SystrayUI) {
	dnsService := dns.DnsService{}
	conectedDnsName := systrayUI.selectedDnsConfiguration.Name
	_, err := dnsService.ChangeDns(dns.SetDns, systrayUI.selectedDnsConfiguration)
	if err != nil {
		fmt.Printf("Error ChangeDns: %v \n", err)
		return
	}

	fmt.Println(conectedDnsName + " set successfully.")

	systrayUI.setIcon(true)
	systrayUI.setToolTip("Connected to: " + conectedDnsName)
}
