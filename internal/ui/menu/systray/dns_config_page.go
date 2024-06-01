package systray

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/ui"
	"github.com/Kingfish219/PlaNet/network/dns"
)

type DnsConfigPage struct {
	systray *SystrayUI
}

func NewDnsConfigPage(systray *SystrayUI) *DnsConfigPage {
	return &DnsConfigPage{
		systray: systray,
	}
}

func (dnsConfig *DnsConfigPage) Initialize() *ui.Page {
	itemList := []ui.Item{
		{
			Key:   "systray_main_dns_config_new",
			Title: "Add New Config",
			Exec: func() {
				//fmt.Print("Hiiii")
			},
			Exec2: func() any {
				vvv := addNewConfig(dnsConfig.systray, "systray_main_dns_config")
				return vvv
			},
		},
	}
	configsList := getExtistedConfig(dnsConfig.systray, "systray_main_dns_config")
	itemList = append(itemList, configsList...)

	return &ui.Page{
		Key:   "systray_main_dns_config",
		Items: itemList,
	}
}

func getExtistedConfig(systrayUI *SystrayUI, key string) []ui.Item {
	configsList := []ui.Item{}
	for _, dnsConfig := range systrayUI.dnsConfigurations {
		exec := func(config dns.Dns) func() {
			return func() {
				DnsConfigOnClick(systrayUI, config, key)
			}
		}(dnsConfig)

		var item = ui.Item{
			Key:   key + "_" + fmt.Sprintf("%v", dnsConfig.Name),
			Title: dnsConfig.Name,
			Exec:  exec,
		}
		configsList = append(configsList, item)

	}
	return configsList
}

func DnsConfigOnClick(systray *SystrayUI, localDns dns.Dns, configKey string) { //, SetTitle func(string)
	systrayUI := *systray
	if systrayUI.connectedDnsConfiguration.Name != localDns.Name {
		dnsService := dns.DnsService{}
		_, err := dnsService.ChangeDns(dns.ResetDns, systrayUI.connectedDnsConfiguration)
		if err != nil {
			fmt.Printf("Error ChangeDns: %v \n", err)
			return
		}
	}
	systrayUI.setIcon(false)
	systrayUI.SystrayMenuItem[configKey].SetTitle(fmt.Sprintf("Config: %v", localDns.Name))
	systrayUI.selectedDnsConfiguration = localDns
}

func addNewConfig(systray *SystrayUI, key string) *ui.Item {
	newDns := dns.Dns{Name: "MyTest", PrimaryDns: "185.51.200.2", SecendaryDns: "178.22.122.100"}
	// newDns := openCMDAndGetDNSData()
	if systray.SystrayMenuItem[newDns.Name] != nil {
		fmt.Println(newDns.Name + " existed")
		return &ui.Item{}
	}

	err := systray.dnsRepository.ModifyDnsConfigurations(newDns)
	if err != nil {
		fmt.Printf("Error : %v \n", err)
		return &ui.Item{}
	}
	systray.selectedDnsConfiguration = newDns

	dnsService := dns.DnsService{}
	_, err = dnsService.ChangeDns(dns.SetDns, systray.connectedDnsConfiguration)
	if err != nil {
		fmt.Printf("Error : %v \n", err)
		return &ui.Item{}
	}

	fmt.Println(newDns.Name + " connected successfully.")

	systray.setIcon(true)
	systray.setToolTip("connected to : " + newDns.Name)

	exec := func(config dns.Dns) func() {
		return func() {
			DnsConfigOnClick(systray, config, key)
		}
	}(newDns)

	var item = &ui.Item{
		Key:   key + "_" + fmt.Sprintf("%v", newDns.Name),
		Title: newDns.Name,
		Exec:  exec,
	}
	return item
	// dnsConfigSubMenu := systray.SystrayMenuItem[configMenuKey].AddSubMenuItem(newDns.Name, newDns.Name)

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
