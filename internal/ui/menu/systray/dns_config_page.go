package systray

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/ui"
	"github.com/Kingfish219/PlaNet/network/dns"

	"github.com/Kingfish219/PlaNet/internal/publisher"
	"github.com/Kingfish219/PlaNet/internal/ui/console"
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
				addNewConfig(dnsConfig.systray)
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
	if systray.connectedDnsConfiguration.Name != localDns.Name {
		dnsService := dns.DnsService{}
		_, err := dnsService.ChangeDns(dns.ResetDns, systray.connectedDnsConfiguration)
		if err != nil {
			fmt.Printf("Error ChangeDns: %v \n", err)
			return
		}
	}
	systray.setIcon(false)
	systray.SystrayMenuItem[configKey].SetTitle(fmt.Sprintf("Config: %v", localDns.Name))
	systray.selectedDnsConfiguration = localDns
}

func addNewConfig(systray *SystrayUI) {

	console := console.New(systray.dnsRepository)
	publisher := publisher.Publisher{}
	publisher.UISubscribers = append(publisher.UISubscribers, console)
	publisher.PublishUI("new-config")
}
