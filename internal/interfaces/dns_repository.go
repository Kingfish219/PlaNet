package interfaces

import "github.com/Kingfish219/PlaNet/network/dns"

type DnsRepository interface {
	GetDnsConfigurations() ([]dns.Dns, error)
	ModifyDnsConfigurations(dns dns.Dns) (bool, error)
}
