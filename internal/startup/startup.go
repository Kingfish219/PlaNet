package startup

import (
	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/io"
)

type Startup struct {
	interactionables []interfaces.Interactionable
}

func New() Startup {
	return Startup{
		interactionables: []interfaces.Interactionable{},
	}
}

func (startup *Startup) Prepare() error {
	console := io.ConsoleIO{}
	startup.interactionables = append(startup.interactionables, console)

	systray := io.SystrayIO{}
	startup.interactionables = append(startup.interactionables, systray)

	return nil
}

func (startup *Startup) Initialize() error {
	var err error

	for _, Interactionable := range startup.interactionables {
		err = Interactionable.Initialize()
	}

	return err
}
