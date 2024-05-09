package systray

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Kingfish219/PlaNet/internal/interfaces"
	"github.com/Kingfish219/PlaNet/internal/presets"
	"github.com/Kingfish219/PlaNet/network/dns"
	"github.com/Kingfish219/PlaNet/network/ni"
	"github.com/getlantern/systray"
)

type SystrayUI struct {
	dnsRepository             interfaces.DnsRepository
	dnsConfigurations         []dns.Dns
	selectedDnsConfiguration  dns.Dns
	connectedDnsConfiguration dns.Dns
}

func New(dnsRepository interfaces.DnsRepository) *SystrayUI {
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

	dnsMenus := map[string]*systray.MenuItem{}

	dnsConfigMenu := systray.AddMenuItem(fmt.Sprintf("DNS config: %v", systrayUI.dnsConfigurations[0].Name), "Selected DNS Configuration")
	for _, dnsConfig := range systrayUI.dnsConfigurations {

		dnsConfigSubMenu := dnsConfigMenu.AddSubMenuItem(dnsConfig.Name, dnsConfig.Name)
		localDns := dnsConfig
		dnsMenus[dnsConfig.Name] = dnsConfigSubMenu
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

	menuSet := systray.AddMenuItem("Set DNS", "Set DNS")
	menuReset := systray.AddMenuItem("Reset DNS", "Reset DNS")
	menuAdd := systray.AddMenuItem("Add DNS", "Add DNS")
	menuDelete := systray.AddMenuItem("Delete Current DNS", "Delete Current DNS")

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
			fmt.Println("dnsMenu", dnsMenus)
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
			<-menuAdd.ClickedCh
			fmt.Println("Add new dns")

			// newDns := dns.Dns{Name: "MyTest", PrimaryDns: "185.51.200.2", SecendaryDns: "178.22.122.100"}
			newDns := openCMDAndGetDNSData()

			if dnsMenus[newDns.Name] != nil {
				fmt.Println(newDns.Name + " existed")
				return
			}

			err = systrayUI.dnsRepository.ModifyDnsConfigurations(newDns)
			if err != nil {
				fmt.Println(err)
				return
			}
			systrayUI.selectedDnsConfiguration = newDns

			dnsService := dns.DnsService{}
			_, err := dnsService.ChangeDns(dns.SetDns, systrayUI.connectedDnsConfiguration)
			if err != nil {
				fmt.Println(err)

				return
			}

			fmt.Println(newDns.Name + " connected successfully.")

			systrayUI.setIcon(true)
			systrayUI.setToolTip("connected to : " + newDns.Name)

			dnsConfigSubMenu := dnsConfigMenu.AddSubMenuItem(newDns.Name, newDns.Name)

			dnsMenus[newDns.Name] = dnsConfigSubMenu
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

			for key, item := range dnsMenus {
				if key == targetDns.Name {
					item.Hide()
					dnsConfigMenu.SetTitle("DNS config: -")
				}
			}
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

	return nil
}

func openCMDAndGetDNSData() dns.Dns {
	cmdPath := filepath.Join(os.Getenv("SystemRoot"), "System32", "cmd.exe")
	pr, pw := io.Pipe()
	cmd := exec.Command(cmdPath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = pr
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return dns.Dns{}
	}

	defer pw.Close()

	fmt.Print("Enter DNS Name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dnsName := scanner.Text()

	pw.Write([]byte(dnsName + "\n"))

	fmt.Print("Enter Primary IP:")
	scanner.Scan()
	primaryDns := scanner.Text()

	// Write the second data to the pipe
	pw.Write([]byte(primaryDns + "\n"))

	// Prompt the user for the third data
	fmt.Print("Enter Secendary Dns:")
	scanner.Scan()
	secendaryDns := scanner.Text()

	// Write the third data to the pipe
	pw.Write([]byte(secendaryDns + "\n"))

	// Close the write end of the pipe
	pw.Close()

	newDns := dns.Dns{Name: dnsName, PrimaryDns: primaryDns, SecendaryDns: secendaryDns}
	return newDns
}
