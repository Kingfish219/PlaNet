package presets

import "github.com/Kingfish219/PlaNet/network/dns"

func GetDnsConfigurations() []dns.Dns {
	dnsList := []dns.Dns{
		{
			Name:         "Shecan",
			PrimaryDns:   "185.51.200.2",
			SecendaryDns: "178.22.122.100",
		},
	}

	return dnsList
}
