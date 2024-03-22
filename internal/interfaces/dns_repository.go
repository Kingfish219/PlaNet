package interfaces

import "github.com/Kingfish219/PlaNet/internal/domain"

type DnsRepository interface {
	GetDnsConfigurations() ([]domain.Dns, error)
	ModifyDnsConfigurations(dns domain.Dns) (bool, error)
}
