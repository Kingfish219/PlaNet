package main

import (
	"PlaNet/src/domain"
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/getlantern/systray"
)

var currentIpConfig domain.IPConfiguration

func main() {
	systray.Run(onReady, onExit)
}

func onExit() {
	fmt.Println("Exiting")
}

func onReady() {
	setIcon(false)
	setToolTip("Not connected")

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
		_, err := changeDns("set", shecan)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Shecan set successfully.")

		setIcon(true)
		setToolTip("Connected to: Shecan")
	}()

	go func() {
		<-menuReset.ClickedCh
		_, err := changeDns("reset", shecan)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Shecan disconnected successfully.")

		setIcon(false)
		setToolTip("Not connected")
	}()
}

func setIcon(status bool) {
	var fileName = "idle"
	if status {
		fileName = "success"
	}

	var filePath = "./assets/" + fileName + ".ico"
	ico, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Unable to read icon:", err)
	} else {
		systray.SetIcon(ico)
	}
}

func setToolTip(toolTip string) {
	systray.SetTooltip("PlaNet:\n" + toolTip)
}

func changeDns(operation string, dns domain.Dns) (bool, error) {
	var activeInterfaceName = getActiveNetworkInterface()
	if activeInterfaceName == "" {
		return false, errors.New("failed to get active network interface")
	}

	if operation == "set" {
		currentIpConfig.IPAddress = ""
		currentIpConfig.SubnetMask = ""
		currentIpConfig.DefaultGateway = ""

		result, ipConfig, _ := getStaticIPConfiguration(activeInterfaceName)
		if result && ipConfig != nil {
			fmt.Println(ipConfig)
			currentIpConfig = *ipConfig
		}

		return setDns(activeInterfaceName, dns.PrimaryDns, dns.SecendaryDns)
	} else {
		var result, err = resetDns(activeInterfaceName)
		if result && currentIpConfig.IPAddress != "" {
			fmt.Println("Reset")
			fmt.Println(currentIpConfig)
			setStaticIPConfiguration(activeInterfaceName, &currentIpConfig)
		}

		return result, err
	}
}

func resetDns(interfaceName string) (bool, error) {
	cmd := exec.Command("netsh", "interface", "ipv4", "set", "dnsservers", "name="+interfaceName, "source=dhcp")
	err := cmd.Run()
	if err != nil {
		return false, err
	}

	return true, nil
}

func setDns(interfaceName string, dns1 string, dns2 string) (bool, error) {
	commandText := fmt.Sprintf(`Set-DNSClientServerAddress -InterfaceAlias "%s" -ServerAddresses ("%s","%s")`,
		interfaceName, dns1, dns2)

	cmd := exec.Command("powershell", "-Command", commandText)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("failed to execute command: %s, output: %s", err, output)
	}

	return true, nil
}

func getStaticIPConfiguration(interfaceName string) (bool, *domain.IPConfiguration, error) {
	cmd := exec.Command("netsh", "interface", "ipv4", "show", "config", "name="+interfaceName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	ipConfig := &domain.IPConfiguration{}
	read := false

	for scanner.Scan() {
		line := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if strings.Contains(line, strings.ToLower("Configuration for interface")) {
			read = true
			continue
		}

		if read && strings.HasPrefix(line, strings.ToLower("DHCP enabled")) {
			fields := strings.Fields(line)
			var enabled = fields[2]
			if strings.EqualFold(enabled, "yes") {
				return true, nil, nil
			}
		}

		if read && strings.HasPrefix(line, strings.ToLower("IP Address")) {
			fields := strings.Fields(line)
			ipConfig.IPAddress = fields[2]
		}

		if read && strings.HasPrefix(line, strings.ToLower("Subnet")) {
			fields := strings.Fields(line)
			ipConfig.SubnetMask = strings.TrimRight(fields[4], ")")
		}

		if read && strings.HasPrefix(line, strings.ToLower("Default Gateway")) {
			fields := strings.Fields(line)
			ipConfig.DefaultGateway = fields[2]
		}
	}

	if ipConfig.IPAddress == "" || ipConfig.SubnetMask == "" {
		return false, ipConfig, fmt.Errorf("no static IP configuration found for interface: %s", interfaceName)
	}

	return true, ipConfig, nil
}

func setStaticIPConfiguration(interfaceName string, ipConfig *domain.IPConfiguration) error {
	cmd := exec.Command("netsh", "interface", "ipv4", "set", "address", "name="+interfaceName,
		"source=static", "addr="+ipConfig.IPAddress, "mask="+ipConfig.SubnetMask,
		"gateway="+ipConfig.DefaultGateway)
	return cmd.Run()
}

// func getOperationFromUser() string {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("What can I do for you?", "\n", "you can either execute 'set' or 'reset'")
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("An error occurred while reading input. Please try again", err)
// 		return ""
// 	}

// 	input = strings.TrimSpace(input)

// 	return input
// }

func getActiveNetworkInterface() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
		return ""
	}

	for _, intf := range interfaces {
		if intf.Flags&net.FlagUp != 0 && !strings.Contains(intf.Flags.String(), "loopback") {
			return intf.Name
		}
	}

	return ""
}
