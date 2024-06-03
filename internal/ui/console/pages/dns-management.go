package pages

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/common"
	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/ui"
	"github.com/Kingfish219/PlaNet/network/dns"
)

func DnsManagement(parent *ui.Page, repo interfaces.DnsRepository) *ui.Page {
	selectedDnsConfig, err := repo.GetSelectedDnsConfiguration()
	if err != nil {
		return &ui.Page{}
	}

	planetConf := common.GetConfigurations()

	page := &ui.Page{
		Key:    "c_dns",
		Title:  "DNS Management",
		Parent: parent,
		Items: []ui.Item{
			{
				Key:      "c_dns_config",
				Title:    fmt.Sprintf("1. Config: %v", selectedDnsConfig.Name),
				ShortKey: "1",
			},
			{
				Key:      "c_dns_set",
				Title:    "2. Set",
				ShortKey: "2",
				Exec: func() any {
					dnsService := dns.DnsService{}
					_, err := dnsService.ChangeDns(dns.SetDns, selectedDnsConfig)
					if err == nil {
						planetConf.ActiveDns = &selectedDnsConfig
					}
					return nil
				},
			},
			{
				Key:      "c_dns_reset",
				Title:    "3. Reset",
				ShortKey: "3",
				Exec: func() any {
					dnsService := dns.DnsService{}
					_, err := dnsService.ChangeDns(dns.ResetDns, selectedDnsConfig)
					if err == nil {
						planetConf.ActiveDns = nil
					}
					return nil
				},
			},
			{
				Key:      "c_dns_delete",
				Title:    fmt.Sprintf("4. Delete selected config: %v", selectedDnsConfig.Name),
				ShortKey: "4",
			},
		},
	}

	page.Items[0].Page = DnsConfig(page, repo)

	return page
}
