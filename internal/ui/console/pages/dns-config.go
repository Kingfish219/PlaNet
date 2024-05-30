package pages

import (
	"fmt"
	"strings"

	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/ui"
)

func DnsConfig(repo interfaces.DnsRepository) *ui.Page {
	dnsConfigurations, err := repo.GetDnsConfigurations()
	if err != nil {
		fmt.Println(err)

		return &ui.Page{}
	}

	items := []ui.Item{}
	for index, dns := range dnsConfigurations {
		items = append(items, ui.Item{
			Key:   fmt.Sprintf("c_dnsconfig_%v", strings.ReplaceAll(dns.Name, " ", "")),
			Title: fmt.Sprintf("%v. %s (%s, %s)", index+1, dns.Name, dns.PrimaryDns, dns.SecendaryDns),
			Exec: func() {
				repo.ModifyActiveDnsConfiguration(dns)
			},
		})
	}

	return &ui.Page{
		Key:   "c_dnsconfig",
		Title: "Select a config:",
		Items: items,
	}
}
