package systray

import (
	"fmt"
	"os"

	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/presets"
	"github.com/Kingfish219/PlaNet/internal/ui"
	"github.com/Kingfish219/PlaNet/network/dns"
	"github.com/Kingfish219/PlaNet/network/ni"
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
	// systrayUI.addDnsConfigurations()
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

func (systrayUI *SystrayUI) addDnsConfigurations() error {
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

	dnsMenu := systray.AddMenuItem("DNS", "DNS")
	dnsConfigMenu := dnsMenu.AddSubMenuItem(fmt.Sprintf("Config: %v", systrayUI.dnsConfigurations[0].Name), "Selected DNS Configuration")

	for _, dnsConfig := range systrayUI.dnsConfigurations {

		dnsConfigSubMenu := dnsConfigMenu.AddSubMenuItem(dnsConfig.Name, dnsConfig.Name)
		localDns := dnsConfig
		systrayUI.SystrayMenuItem[dnsConfig.Name] = dnsConfigSubMenu
		go func(localDns dns.Dns) {
			for {
				<-dnsConfigSubMenu.ClickedCh
				if systrayUI.connectedDnsConfiguration.Name != localDns.Name {
					dnsService := dns.DnsService{}
					_, err := dnsService.ChangeDns(dns.ResetDns, systrayUI.connectedDnsConfiguration)
					if err != nil {
						fmt.Println(err)

						return
					}
				}

				systrayUI.setIcon(false)
				dnsConfigMenu.SetTitle(fmt.Sprintf("DNS config: %v", localDns.Name))
				systrayUI.selectedDnsConfiguration = localDns
			}

		}(localDns)
	}

	menuAdd := dnsConfigMenu.AddSubMenuItem("Add New Config", "Add New Config")

	menuSet := dnsMenu.AddSubMenuItem("Set", "Set DNS")
	menuReset := dnsMenu.AddSubMenuItem("Reset", "Reset DNS")
	menuDelete := dnsMenu.AddSubMenuItem("Delete This Config", "Delete This Config")

	networkInterfaceMenu := systray.AddMenuItem("Network Interface", "Network Interface")
	toolsMenu := systray.AddMenuItem("Tools", "Tools")
	consoleMenu := systray.AddMenuItem("Console", "Console")

	go func() {
		for {
			<-menuAdd.ClickedCh
			fmt.Println("Add new dns")

			newDns := dns.Dns{Name: "MyTest", PrimaryDns: "185.51.200.2", SecendaryDns: "178.22.122.100"}
			// newDns := openCMDAndGetDNSData()

			if systrayUI.SystrayMenuItem[newDns.Name] != nil {
				fmt.Println(newDns.Name + " existed")
				return
			}

			err := systrayUI.dnsRepository.ModifyDnsConfigurations(newDns)
			if err != nil {
				fmt.Println(err)
				return
			}
			systrayUI.selectedDnsConfiguration = newDns

			dnsService := dns.DnsService{}
			_, err = dnsService.ChangeDns(dns.SetDns, systrayUI.connectedDnsConfiguration)
			if err != nil {
				fmt.Println(err)

				return
			}

			fmt.Println(newDns.Name + " connected successfully.")

			systrayUI.setIcon(true)
			systrayUI.setToolTip("connected to : " + newDns.Name)

			dnsConfigSubMenu := dnsConfigMenu.AddSubMenuItem(newDns.Name, newDns.Name)

			systrayUI.SystrayMenuItem[newDns.Name] = dnsConfigSubMenu
			localDns := newDns
			dnsConfigMenu.SetTitle(fmt.Sprintf("DNS config: %v", newDns.Name))

			go func(localDns dns.Dns) {
				for {
					<-dnsConfigSubMenu.ClickedCh
					if systrayUI.connectedDnsConfiguration.Name != localDns.Name {
						dnsService := dns.DnsService{}
						_, err := dnsService.ChangeDns(dns.ResetDns, systrayUI.connectedDnsConfiguration)
						if err != nil {
							fmt.Println(err)
							return
						}
					}

					systrayUI.setIcon(false)
					dnsConfigMenu.SetTitle(fmt.Sprintf("DNS config: %v", localDns.Name))
					systrayUI.selectedDnsConfiguration = localDns
				}

			}(localDns)
		}

	}()

	go func() {
		for {
			<-menuSet.ClickedCh
			dnsService := dns.DnsService{}
			_, err := dnsService.ChangeDns(dns.SetDns, systrayUI.selectedDnsConfiguration)
			if err != nil {
				fmt.Println(err)

				return
			}

			fmt.Println("Shecan set successfully.")

			systrayUI.setIcon(true)
			systrayUI.setToolTip("Connected to: Shecan")
		}

	}()

	go func() {
		for {
			<-menuReset.ClickedCh
			fmt.Println(systrayUI.selectedDnsConfiguration)
			fmt.Println("dnsMenu", systrayUI.SystrayMenuItem)
			dnsService := dns.DnsService{}
			_, err := dnsService.ChangeDns(dns.ResetDns, systrayUI.connectedDnsConfiguration)
			if err != nil {
				fmt.Println(err)

				return
			}

			fmt.Println("Shecan disconnected successfully.")

			systrayUI.setIcon(false)
			systrayUI.setToolTip("Not connected")
		}

	}()

	go func() {
		for {
			<-menuDelete.ClickedCh
			fmt.Println("Delete Current Dns")

			activeInterfaceNames, err := ni.GetActiveNetworkInterface()
			if activeInterfaceNames == nil || err != nil {
				fmt.Println(err)

				return
			}

			if len(activeInterfaceNames) == 0 {
				fmt.Println("no active network interface found")

				return
			}

			var targetDns = systrayUI.selectedDnsConfiguration
			err = systrayUI.dnsRepository.DeleteDnsConfigurations(systrayUI.selectedDnsConfiguration)
			if err != nil {
				fmt.Println(err)
				return
			}

			for key, item := range systrayUI.SystrayMenuItem {
				if key == targetDns.Name {
					item.Hide()
					dnsConfigMenu.SetTitle("DNS config: -")
				}
			}
			delete(systrayUI.SystrayMenuItem, targetDns.Name)
			dnsService := dns.DnsService{}
			_, err = dnsService.ChangeDns(dns.ResetDns, systrayUI.connectedDnsConfiguration)
			if err != nil {
				fmt.Println(err)

				return
			}

			fmt.Println(targetDns.Name + " deleted successfully.")

			systrayUI.setIcon(false)
			systrayUI.setToolTip("Not connected")
		}

	}()

	go func() {
		for {
			<-networkInterfaceMenu.ClickedCh
			fmt.Println("Network Interface")

		}

	}()

	go func() {
		for {
			<-toolsMenu.ClickedCh
			fmt.Println("Tools")

		}

	}()

	go func() {
		for {
			<-consoleMenu.ClickedCh
			fmt.Println("console")

		}

	}()

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
					if pageItem.Key == "systray_main_dns_config_new" && pageItem.Exec2 != nil {
						returnVal := pageItem.Exec2()
						addSystrayMenu(returnVal.(*ui.Item), systrayUI.SystrayMenuItem["systray_main_dns_config"], systrayUI)
					}
				}

			}(pageItem)
		}

	}
}
