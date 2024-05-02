package presets

import "github.com/Kingfish219/PlaNet/network/dns"

func GetDnsPresets() []dns.Dns {
	dnsList := []dns.Dns{
		{
			Name:         "Google",
			PrimaryDns:   "8.8.8.8",
			SecendaryDns: "8.8.4.4",
		},
		{
			Name:         "Shecan",
			PrimaryDns:   "185.51.200.2",
			SecendaryDns: "178.22.122.100",
		},
		{
			Name:         "CloudFlare",
			PrimaryDns:   "1.1.1.1",
			SecendaryDns: "1.0.0.1",
		},
	}

	return dnsList
}
