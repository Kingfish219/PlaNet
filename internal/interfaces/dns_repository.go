package interfaces

import "github.com/Kingfish219/PlaNet/network/dns"

type DnsRepository interface {
	GetSelectedDnsConfiguration() (dns.Dns, error)
	ModifyActiveDnsConfiguration(dns.Dns) error
	GetDnsConfigurations() ([]dns.Dns, error)
	ModifyDnsConfigurations(dns.Dns) error
}
