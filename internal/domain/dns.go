package domain

type Dns struct {
	Name         string `json:"name"`
	PrimaryDns   string `json:"primary_dns"`
	SecendaryDns string `json:"secondary_dns"`
}
