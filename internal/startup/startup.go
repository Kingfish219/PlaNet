package startup

import (
	"fmt"
	"os"

	"github.com/Kingfish219/PlaNet/internal/interfaces"
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
	fmt.Println(os.UserConfigDir())
	fmt.Println(os.UserHomeDir())
	fmt.Println(os.TempDir())
	fmt.Println(os.UserCacheDir())

	// dnsRepository := repository.NewDnsRepository("")

	// console := ui.ConsoleUI{}
	// startup.userInterfaces = append(startup.userInterfaces, console)

	// systray := ui.NewSystrayUI(dnsRepository)
	// startup.userInterfaces = append(startup.userInterfaces, systray)

	return nil
}

func (startup *Startup) Start() error {
	var err error

	for _, userInterface := range startup.userInterfaces {
		err = userInterface.Initialize()
	}

	return err
}
