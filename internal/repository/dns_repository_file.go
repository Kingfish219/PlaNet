package repository

import (
	"encoding/json"
	"os"

	"github.com/Kingfish219/PlaNet/internal/domain"
)

type DnsRepositoryFile struct {
	filePath string
}

func NewDnsRepository(filePath string) *DnsRepositoryFile {
	return &DnsRepositoryFile{
		filePath: filePath,
	}
}

func (repo DnsRepositoryFile) GetDnsConfigurations() ([]domain.Dns, error) {
	file, err := os.ReadFile(repo.filePath)
	if err != nil {
		return nil, err
	}

	var dnsList []domain.Dns
	err = json.Unmarshal(file, &dnsList)
	if err != nil {
		return nil, err
	}

	return dnsList, nil
}

func (repo DnsRepositoryFile) ModifyDnsConfigurations(dns domain.Dns) (bool, error) {
	dnsList, err := repo.GetDnsConfigurations()
	if err != nil {
		return false, err
	}

	dnsList = append(dnsList, dns)
	var jsonData []byte

	jsonData, err = json.Marshal(dnsList)
	if err != nil {
		return false, err
	}

	err = os.WriteFile(repo.filePath, jsonData, 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}
