package presets

import "github.com/Kingfish219/PlaNet/internal/domain"

func GetDnsConfigurations() []domain.Dns {
	dnsList := []domain.Dns{
		{
			Name:         "Shecan",
			PrimaryDns:   "185.51.200.2",
			SecendaryDns: "178.22.122.100",
		},
	}

	return dnsList
}
