package pages

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/ui"
)

func NewDns(parent *ui.Page, repo interfaces.DnsRepository) *ui.Page {
	page := &ui.Page{
		Key:    "c_dnsconfig",
		Title:  "Add a new DNS config:",
		Parent: parent,
	}

	page.Initiate = func() {
		fmt.Println()
		fmt.Println("Primary Dns: ")
	}

	return page
}
