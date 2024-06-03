package pages

import (
	"fmt"
	"strings"

	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/ui"
)

func DnsConfig(parent *ui.Page, repo interfaces.DnsRepository) *ui.Page {
	dnsConfigurations, err := repo.GetDnsConfigurations()
	if err != nil {
		fmt.Println(err)

		return &ui.Page{}
	}

	items := []ui.Item{}
	for index, dns := range dnsConfigurations {
		items = append(items, ui.Item{
			Key:      fmt.Sprintf("c_dnsconfig_%v", strings.ReplaceAll(dns.Name, " ", "")),
			Title:    fmt.Sprintf("%v. %s (%s, %s)", index+1, dns.Name, dns.PrimaryDns, dns.SecendaryDns),
			ShortKey: fmt.Sprint(index + 1),
			Exec: func() any {
				repo.ModifyActiveDnsConfiguration(dns)
				return nil
			},
		})
	}

	page := &ui.Page{
		Key:    "c_dnsconfig",
		Title:  "Select a config:",
		Parent: parent,
		Items:  items,
	}

	title := fmt.Sprintf("%v. Add a new config", len(items)+1)
	page.Items = append(items, ui.Item{
		Key:      "c_dnsconfig_addnewconfig",
		Title:    title,
		ShortKey: fmt.Sprint(len(page.Items) + 1),
		Page:     NewDns(page, repo),
	})

	return page
}
