package pages

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/ui"
	"github.com/Kingfish219/PlaNet/network/dns"
)

func DnsManagement(repo interfaces.DnsRepository) *ui.Page {
	activeDnsConfig, err := repo.GetActiveDnsConfiguration()
	if err != nil {
		return &ui.Page{}
	}

	return &ui.Page{
		Key:   "c_dns",
		Title: "DNS Management",
		Items: []ui.Item{
			{
				Key:      "c_dns_config",
				Title:    fmt.Sprintf("1. Config: %v", activeDnsConfig.Name),
				ShortKey: "1",
				Page:     DnsConfig(repo),
			},
			{
				Key:      "c_dns_set",
				Title:    "2. Set",
				ShortKey: "2",
				Exec: func() any {
					dnsService := dns.DnsService{}
					dnsService.ChangeDns(dns.SetDns, activeDnsConfig)
					return nil
				},
			},
			{
				Key:      "c_dns_reset",
				Title:    "3. Reset",
				ShortKey: "3",
				Exec: func() any {
					dnsService := dns.DnsService{}
					dnsService.ChangeDns(dns.ResetDns, activeDnsConfig)
					return nil
				},
			},
			{
				Key:      "c_dns_delete",
				Title:    fmt.Sprintf("4. Delete selected config: %v", activeDnsConfig.Name),
				ShortKey: "4",
			},
		},
	}
}
