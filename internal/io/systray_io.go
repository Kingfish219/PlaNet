package io

import (
	"fmt"
	"os"

	"github.com/Kingfish219/PlaNet/internal/domain"
	"github.com/Kingfish219/PlaNet/network"
	"github.com/getlantern/systray"
)

type SystrayIO struct {
}

func (iSystray SystrayIO) Initialize() error {
	systray.Run(iSystray.onReady, iSystray.onExit)

	return nil
}

func (iSystray SystrayIO) onExit() {
	fmt.Println("Exiting")
}

func (iSystray SystrayIO) onReady() {
	iSystray.setIcon(false)
	iSystray.setToolTip("Not connected")
	iSystray.addDnsConfigurations()

	menuSet := systray.AddMenuItem("Set", "Set DNS")
	menuReset := systray.AddMenuItem("Reset", "Reset DNS")
	menuExit := systray.AddMenuItem("Exit", "Exit the application")

	var shecan = domain.Dns{
		Name:         "Shecan",
		PrimaryDns:   "185.51.200.2",
		SecendaryDns: "178.22.122.100",
	}

	go func() {
		<-menuExit.ClickedCh
		systray.Quit()
	}()

	go func() {
		<-menuSet.ClickedCh
		dnsService := network.DnsService{}
		_, err := dnsService.ChangeDns("set", shecan)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("Shecan set successfully.")

		iSystray.setIcon(true)
		iSystray.setToolTip("Connected to: Shecan")
	}()

	go func() {
		<-menuReset.ClickedCh
		dnsService := network.DnsService{}
		_, err := dnsService.ChangeDns("reset", shecan)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("Shecan disconnected successfully.")

		iSystray.setIcon(false)
		iSystray.setToolTip("Not connected")
	}()
}

func (iSystray SystrayIO) setIcon(status bool) {
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

func (iSystray SystrayIO) setToolTip(toolTip string) {
	systray.SetTooltip("PlaNet:\n" + toolTip)
}

func (iSystray SystrayIO) addDnsConfigurations() {

}
