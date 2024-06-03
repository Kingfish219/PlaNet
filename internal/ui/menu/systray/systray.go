package systray

import (
	"fmt"
	"os"

	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/presets"
	"github.com/Kingfish219/PlaNet/internal/ui"
	"github.com/Kingfish219/PlaNet/network/dns"
	"github.com/getlantern/systray"
)

type SystrayUI struct {
	dnsRepository             interfaces.DnsRepository
	dnsConfigurations         []dns.Dns
	selectedDnsConfiguration  dns.Dns
	connectedDnsConfiguration dns.Dns
	Page                      *ui.Page
	SystrayMenuItem           map[string]*systray.MenuItem
}

func New(dnsRepository interfaces.DnsRepository) *SystrayUI {

	return &SystrayUI{
		dnsRepository:   dnsRepository,
		SystrayMenuItem: map[string]*systray.MenuItem{},
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

	systrayUI.addMenu()

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

func (systray *SystrayUI) Consume(command string) error {
	switch command {
	case "new-config":
		systray.refreshConfigsMenu()
		fmt.Println("systray new config")
		break
	}
	return nil
}

func (systrayUI *SystrayUI) addMenu() error {
	dnsConfigurations, err := systrayUI.dnsRepository.GetDnsConfigurations()
	if err != nil {
		return err
	}
	if len(dnsConfigurations) == 0 {
		presetDnsList := presets.GetDnsPresets()
		for _, pre := range presetDnsList {
			systrayUI.dnsRepository.ModifyDnsConfigurations(pre)
		}

		dnsConfigurations, err = systrayUI.dnsRepository.GetDnsConfigurations()
		if err != nil {
			return err
		}
	}

	systrayUI.dnsConfigurations = dnsConfigurations
	err = Feed(systrayUI)
	if err != nil {
		return err
	}
	for _, item := range systrayUI.Page.Items {
		addSystrayMenu(&item, nil, systrayUI)
	}

	return nil
}

func addSystrayMenu(pageItemPtr *ui.Item, parentMenu *systray.MenuItem, systrayUI *SystrayUI) {
	pageItem := *pageItemPtr
	if parentMenu == nil {
		mainMenu := systray.AddMenuItem(pageItem.Title, pageItem.Title)
		systrayUI.SystrayMenuItem[pageItem.Key] = mainMenu
		if pageItem.Exec != nil {
			go func(pageItem ui.Item) {
				for {
					<-mainMenu.ClickedCh
					pageItem.Exec()
				}

			}(pageItem)

		}

		if pageItem.Page != nil && len(pageItem.Page.Items) > 0 {
			for _, item := range pageItem.Page.Items {
				addSystrayMenu(&item, mainMenu, systrayUI)

			}
		}

	} else {
		subMenu := parentMenu.AddSubMenuItem(pageItem.Title, pageItem.Title)
		systrayUI.SystrayMenuItem[pageItem.Key] = subMenu
		if pageItem.Page != nil && len(pageItem.Page.Items) > 0 {
			for _, item := range pageItem.Page.Items {
				addSystrayMenu(&item, subMenu, systrayUI)
			}
		}
		if pageItem.Exec != nil {
			go func(pageItem ui.Item) {
				for {
					<-subMenu.ClickedCh
					pageItem.Exec()
				}

			}(pageItem)
		}

	}
}

func (systrayUI *SystrayUI) refreshConfigsMenu() error {
	dnsConfigurations, err := systrayUI.dnsRepository.GetDnsConfigurations()
	if err != nil {
		return err
	}
	systrayUI.dnsConfigurations = dnsConfigurations

	for _, dnsConfig := range systrayUI.dnsConfigurations {

		if systrayUI.SystrayMenuItem["systray_main_dns_config_"+fmt.Sprintf("%v", dnsConfig.Name)] != nil {
			continue
		}

		exec := func(config dns.Dns) func() {
			return func() {
				DnsConfigOnClick(systrayUI, config, "systray_main_dns_config")
			}
		}(dnsConfig)

		var item = &ui.Item{
			Key:   "systray_main_dns_config_" + fmt.Sprintf("%v", dnsConfig.Name),
			Title: dnsConfig.Name,
			Exec:  exec,
		}

		addSystrayMenu(item, systrayUI.SystrayMenuItem["systray_main_dns_config"], systrayUI)

	}
	return nil
}
