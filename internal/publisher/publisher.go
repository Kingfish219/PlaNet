package publisher

import "github.com/Kingfish219/PlaNet/internal/interfaces"

type Publisher struct {
	UISubscribers []interfaces.UserInterface
}

func (publisher *Publisher) PublishUI(command string) error {
	for _, ui := range publisher.UISubscribers {
		err := ui.Consume(command)
		if err != nil {
			return err
		}
	}

	return nil
}
