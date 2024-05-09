package repository

import (
	"encoding/json"
	"os"

	"github.com/Kingfish219/PlaNet/network/dns"
)

type DnsRepositoryFile struct {
	filePath string
}

func NewDnsRepository(filePath string) *DnsRepositoryFile {
	return &DnsRepositoryFile{
		filePath: filePath,
	}
}

func (repo DnsRepositoryFile) GetDnsConfigurations() ([]dns.Dns, error) {
	file, err := os.ReadFile(repo.filePath)
	if err != nil {
		return nil, err
	}

	if len(file) == 0 {
		return []dns.Dns{}, nil
	}

	var dnsList []dns.Dns
	err = json.Unmarshal(file, &dnsList)
	if err != nil {
		return nil, err
	}

	return dnsList, nil
}

func (repo DnsRepositoryFile) ModifyDnsConfigurations(dns dns.Dns) error {
	dnsList, err := repo.GetDnsConfigurations()
	if err != nil {
		return err
	}

	dnsList = append(dnsList, dns)
	var jsonData []byte

	jsonData, err = json.Marshal(dnsList)
	if err != nil {
		return err
	}

	err = os.WriteFile(repo.filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (repo DnsRepositoryFile) DeleteDnsConfigurations(dns dns.Dns) error {
	dnsList, err := repo.GetDnsConfigurations()
	if err != nil {
		return err
	}

	if targetDnsIndex, err := repo.GetDnsIndex(dns); err == nil {
		dnsList = append(dnsList[:targetDnsIndex], dnsList[targetDnsIndex+1:]...)
	}

	var jsonData []byte

	jsonData, err = json.Marshal(dnsList)
	if err != nil {
		return err
	}

	err = os.WriteFile(repo.filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (repo DnsRepositoryFile) GetDnsIndex(dns dns.Dns) (int, error) {
	dnsList, err := repo.GetDnsConfigurations()
	if err != nil {
		return 0, err
	}
	for index, item := range dnsList {
		if item.Name == dns.Name {
			return index, nil
		}
	}
	return 0, err
}
