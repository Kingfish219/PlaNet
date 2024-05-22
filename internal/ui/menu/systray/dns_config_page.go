package systray

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/ui"
	"github.com/Kingfish219/PlaNet/network/dns"
)

type DnsConfigPage struct {
	systray        *SystrayUI
	key            string
	SetConfigTitle func(string)
}

func NewDnsConfigPage(systray *SystrayUI, key string) *DnsConfigPage {
	return &DnsConfigPage{
		systray: systray,
		key:     key,
	}
}

func (dnsConfig *DnsConfigPage) Initialize() ui.Page {
	itemList := []ui.Item{
		{
			Key:   dnsConfig.key + "_new",
			Title: "Add New Config",
			Exec: func() {
				addNewConfig(dnsConfig.systray)
			},
		},
	}
	configsList := getExtistedConfig(dnsConfig.systray, dnsConfig.key, dnsConfig.SetConfigTitle)
	itemList = append(itemList, configsList...)

	return ui.Page{
		Key:   dnsConfig.key,
		Items: itemList,
	}
}

func getExtistedConfig(systrayUI *SystrayUI, key string, SetConfigTitle func(string)) []ui.Item {
	configsList := []ui.Item{}
	for index, dnsConfig := range systrayUI.dnsConfigurations {
		var item = ui.Item{
			Key:   key + "_" + fmt.Sprintf("%v", index),
			Title: dnsConfig.Name,
			Exec: func() {
				DnsConfigOnClick(systrayUI, dnsConfig, SetConfigTitle)
			},
		}
		configsList = append(configsList, item)

	}
	return configsList
}

func DnsConfigOnClick(systray *SystrayUI, localDns dns.Dns, SetConfigTitle func(string)) { //, SetTitle func(string)
	if systray.connectedDnsConfiguration.Name != localDns.Name {
		dnsService := dns.DnsService{}
		_, err := dnsService.ChangeDns(dns.ResetDns, systray.connectedDnsConfiguration)
		if err != nil {
			fmt.Println(err)

			return
		}
	}
	systray.setIcon(false)
	//	SetConfigTitle(fmt.Sprintf("DNS config: %v", localDns.Name))
	systray.selectedDnsConfiguration = localDns
}

func addNewConfig(systray *SystrayUI) {
	newDns := dns.Dns{Name: "MyTest", PrimaryDns: "185.51.200.2", SecendaryDns: "178.22.122.100"}
	// newDns := openCMDAndGetDNSData()

	if systray.DnsMenus[newDns.Name] != nil {
		fmt.Println(newDns.Name + " existed")
		return
	}

	err := systray.dnsRepository.ModifyDnsConfigurations(newDns)
	if err != nil {
		fmt.Println(err)
		return
	}
	systray.selectedDnsConfiguration = newDns

	dnsService := dns.DnsService{}
	_, err = dnsService.ChangeDns(dns.SetDns, systray.connectedDnsConfiguration)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(newDns.Name + " connected successfully.")

	systray.setIcon(true)
	systray.setToolTip("connected to : " + newDns.Name)

	// dnsConfigSubMenu := dnsConfigMenu.AddSubMenuItem(newDns.Name, newDns.Name)

	// dnsMenus[newDns.Name] = dnsConfigSubMenu
	// localDns := newDns
	// dnsConfigMenu.SetTitle(fmt.Sprintf("DNS config: %v", newDns.Name))

	// go func(localDns dns.Dns) {
	// 	for {
	// 		<-dnsConfigSubMenu.ClickedCh
	// 		if systrayUI.connectedDnsConfiguration.Name != localDns.Name {
	// 			dnsService := dns.DnsService{}
	// 			_, err := dnsService.ChangeDns(dns.ResetDns, systrayUI.connectedDnsConfiguration)
	// 			if err != nil {
	// 				fmt.Println(err)
	// 				return
	// 			}
	// 		}

	// 		systrayUI.setIcon(false)
	// 		dnsConfigMenu.SetTitle(fmt.Sprintf("DNS config: %v", localDns.Name))
	// 		systrayUI.selectedDnsConfiguration = localDns
	// 	}

	// }(localDns)

}
