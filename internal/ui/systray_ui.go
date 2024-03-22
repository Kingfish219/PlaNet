package ui

import (
	"fmt"
	"os"

	"github.com/Kingfish219/PlaNet/internal/domain"
	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/network"
	"github.com/getlantern/systray"
)

type SystrayUI struct {
	dnsRepository     interfaces.DnsRepository
	dnsConfigurations []domain.Dns
}

func NewSystrayUI(dnsRepository interfaces.DnsRepository) *SystrayUI {
	return &SystrayUI{
		dnsRepository: dnsRepository,
	}
}

func (systrayUI *SystrayUI) Initialize() error {
	systray.Run(systrayUI.onReady, systrayUI.onExit)

	return nil
}

func (systrayUI *SystrayUI) onExit() {
	fmt.Println("Exiting")
}

func (systrayUI *SystrayUI) onReady() {
	systrayUI.setIcon(false)
	systrayUI.setToolTip("Not connected")

	systrayUI.addDnsConfigurations()

	systray.AddSeparator()
	menuExit := systray.AddMenuItem("Exit", "Exit the application")

	go func() {
		<-menuExit.ClickedCh
		systray.Quit()
	}()
}

func (systrayUI *SystrayUI) setIcon(status bool) {
	fileName := "idle"
	if status {
		fileName = "success"
	}

	filePath := "./assets/" + fileName + ".ico"
	ico, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Unable to read icon:", err)
	} else {
		systray.SetIcon(ico)
	}
}

func (systrayUI *SystrayUI) setToolTip(toolTip string) {
	systray.SetTooltip("PlaNet:\n" + toolTip)
}

func (systrayUI *SystrayUI) addDnsConfigurations() error {
	dnsConfigurations, err := systrayUI.dnsRepository.GetDnsConfigurations()
	if err != nil {
		return err
	}
	systrayUI.dnsConfigurations = dnsConfigurations

	// var shecan = domain.Dns{
	// 	Name:         "Shecan",
	// 	PrimaryDns:   "185.51.200.2",
	// 	SecendaryDns: "178.22.122.100",
	// }

	menuSet := systray.AddMenuItem("Set", "Set DNS")
	menuReset := systray.AddMenuItem("Reset", "Reset DNS")

	go func() {
		<-menuSet.ClickedCh
		dnsService := network.DnsService{}
		_, err := dnsService.ChangeDns(network.Set, systrayUI.dnsConfigurations[0])
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("Shecan set successfully.")

		systrayUI.setIcon(true)
		systrayUI.setToolTip("Connected to: Shecan")
	}()

	go func() {
		<-menuReset.ClickedCh
		dnsService := network.DnsService{}
		_, err := dnsService.ChangeDns(network.Reset, systrayUI.dnsConfigurations[0])
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("Shecan disconnected successfully.")

		systrayUI.setIcon(false)
		systrayUI.setToolTip("Not connected")
	}()

	return nil
}
