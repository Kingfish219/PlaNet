package pages

import (
	"fmt"
	"strings"

	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/ui"
)

func DnsConfig(repo interfaces.DnsRepository) ui.Page {
	dnsConfigurations, err := repo.GetDnsConfigurations()
	if err != nil {
		return ui.Page{}
	}

	items := []ui.Item{}
	for _, dns := range dnsConfigurations {
		items = append(items, ui.Item{
			Key:   fmt.Sprintf("c_dnsconfig_%v", strings.ReplaceAll(dns.Name, " ", "")),
			Title: dns.Name,
			Exec: func() {
				repo.ModifyActiveDnsConfiguration(dns)
			},
		})
	}

	return ui.Page{
		Key:   "c_dnsconfig",
		Title: "Select ",
		Items: items,
	}
}
