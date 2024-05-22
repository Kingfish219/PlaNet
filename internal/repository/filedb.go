package repository

import "github.com/Kingfish219/PlaNet/network/dns"

type FileDb struct {
	ActiveDns         dns.Dns
	DnsConfigurations []dns.Dns
}
