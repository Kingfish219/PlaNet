package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onExit() {
	fmt.Println("Exiting")
}

func onReady() {
	setIcon(false)
	setToolTip("Idle")

	mSet := systray.AddMenuItem("Set", "Set DNS")
	mReset := systray.AddMenuItem("Reset", "Reset DNS")
	mExit := systray.AddMenuItem("Exit", "Exit the application")

	go func() {
		<-mExit.ClickedCh
		systray.Quit()
	}()

	go func() {
		<-mSet.ClickedCh
		_, err := ChangeDns("set")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Shecan set successfully.")

		setIcon(true)
		setToolTip("Connected to Shecan")
	}()

	go func() {
		<-mReset.ClickedCh
		_, err := ChangeDns("reset")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Shecan disconnected successfully.")

		setIcon(false)
		setToolTip("Idle")
	}()
}

func setIcon(status bool) {
	var fileName = "idle"
	if status {
		fileName = "success"
	}

	var filePath = "./" + fileName + ".ico"
	ico, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Unable to read icon:", err)
	} else {
		systray.SetIcon(ico)
	}
}

func setToolTip(toolTip string) {
	systray.SetTooltip("PlaNet: " + toolTip)
}

func ChangeDns(operation string) (bool, error) {
	var activeInterface = getActiveNetworkInterface()
	if activeInterface == "" {
		return false, errors.New("Failed to get active network interface")
	}

	if operation == "set" {
		return setDns("185.51.200.2", "178.22.122.100")
	} else {
		return resetDns(activeInterface)
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

func setDns(dns1 string, dns2 string) (bool, error) {
	var activeInterface = getActiveNetworkInterface()
	if activeInterface == "" {
		return false, fmt.Errorf("Failed getting active network interface")
	}

	commandText := fmt.Sprintf(`Set-DNSClientServerAddress -InterfaceAlias "%s" -ServerAddresses ("%s","%s")`,
		activeInterface, dns1, dns2)

	cmd := exec.Command("powershell", "-Command", commandText)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("Failed to execute command: %s, Output: %s\n", err, output)
	}

	return true, nil
}

func getOperationFromUser() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What can I do for you?", "\n", "you can either execute 'set' or 'reset'")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return ""
	}

	input = strings.TrimSpace(input)

	return input
}

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
