package ui

import (
	"fmt"
	"os"

	"github.com/Kingfish219/PlaNet/internal/domain"
	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/presets"
	"github.com/Kingfish219/PlaNet/network"
	"github.com/getlantern/systray"
)

type SystrayUI struct {
	dnsRepository            interfaces.DnsRepository
	dnsConfigurations        []domain.Dns
	selectedDnsConfiguration domain.Dns
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

	if len(dnsConfigurations) == 0 {
		presetDnsList := presets.GetDnsConfigurations()
		for _, pre := range presetDnsList {
			systrayUI.dnsRepository.ModifyDnsConfigurations(pre)
		}

		dnsConfigurations, err = systrayUI.dnsRepository.GetDnsConfigurations()
		if err != nil {
			return err
		}
	}

	systrayUI.dnsConfigurations = dnsConfigurations

	var dnsConfigSubMenu *systray.MenuItem

	dnsConfigMenu := systray.AddMenuItem(fmt.Sprintf("DNS config: %v", systrayUI.dnsConfigurations[0].Name), "Selected DNS Configuration")
	for _, dns := range systrayUI.dnsConfigurations {
		dnsConfigSubMenu = dnsConfigMenu.AddSubMenuItem(dns.Name, dns.Name)
	}

	menuSet := systray.AddMenuItem("Set DNS", "Set DNS")
	menuReset := systray.AddMenuItem("Reset DNS", "Reset DNS")

	go func() {
		<-dnsConfigSubMenu.ClickedCh
		fmt.Println(dnsConfigSubMenu.String())
	}()

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
