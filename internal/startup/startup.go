package startup

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/presets"
	"github.com/Kingfish219/PlaNet/internal/repository"
	"github.com/Kingfish219/PlaNet/internal/ui/console"
	"github.com/Kingfish219/PlaNet/internal/ui/menu/systray"
)

type Startup struct {
	userInterfaces []interfaces.UserInterface
}

func New() Startup {
	return Startup{
		userInterfaces: []interfaces.UserInterface{},
	}
}

func (startup *Startup) Initialize() error {
	repoFilePath, err := startup.createRepoFilePath()
	if err != nil {
		return err
	}

	dnsRepository := repository.NewDnsRepository(repoFilePath)
	err = startup.migrateDb(dnsRepository)
	if err != nil {
		return err
	}

	systray := systray.New(dnsRepository)
	startup.userInterfaces = append(startup.userInterfaces, systray)

	console := console.New(dnsRepository)
	startup.userInterfaces = append(startup.userInterfaces, console)

	return nil
}

func (startup *Startup) Start() error {
	var err error

	for _, userInterface := range startup.userInterfaces {
		err = userInterface.Initialize()
	}

	return err
}

func (startup *Startup) createRepoFilePath() (string, error) {
	tempDirPath, err := os.UserCacheDir()
	if err != nil {
		tempDirPath = os.TempDir()
	}

	planetTempDirPath := filepath.Join(tempDirPath, "PlaNet")
	err = os.MkdirAll(planetTempDirPath, 0644)
	if err != nil {
		return "", err
	}

	repoFilePath := filepath.Join(planetTempDirPath, "config.json")
	_, err = os.Stat(repoFilePath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return "", nil
		}

		_, err = os.Create(repoFilePath)
		if err != nil {
			return "", err
		}
	}

	return repoFilePath, nil
}

func (startup *Startup) migrateDb(repository interfaces.DnsRepository) error {
	dnsConfigurations, err := repository.GetDnsConfigurations()
	if err != nil {
		return err
	}
	if len(dnsConfigurations) == 0 {
		presetDnsList := presets.GetDnsPresets()
		for _, pre := range presetDnsList {
			repository.ModifyDnsConfigurations(pre)
		}
	}
	return nil
}
