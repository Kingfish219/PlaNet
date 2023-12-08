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

	// for {
	// 	var operation = getOperationFromUser()

	// 	if operation == "set" || operation == "reset" {
	// 		result, err := ChangeDns(operation)
	// 		if err != nil {
	// 			fmt.Fprintln(os.Stderr, "Error while setting DNS:", err)
	// 			continue
	// 		}

	// 		println(result)
	// 	} else {
	// 		println("Invalid operation. try 'set' or 'reset'")
	// 	}
	// }
}

func onExit() {
	fmt.Println("Exiting")
}

func onReady() {
	// ico, err := os.ReadFile("path/to/icon.ico")
	// if err != nil {
	// 	fmt.Println("Unable to read icon:", err)
	// 	return
	// }
	// systray.SetIcon(ico)

	systray.SetTooltip("PlaNet: manage your system DNS")

	mExit := systray.AddMenuItem("Exit", "Exit the application")
	mSet := systray.AddMenuItem("Set", "Set DNS")
	mReset := systray.AddMenuItem("Reset", "Reset DNS")

	go func() {
		<-mExit.ClickedCh
		systray.Quit()
		fmt.Println("Quit the app")
	}()

	go func() {
		<-mSet.ClickedCh
		ChangeDns("set")
	}()

	go func() {
		<-mReset.ClickedCh
		ChangeDns("reset")
	}()
}

func ChangeDns(operation string) (string, error) {
	var activeInterface = getActiveNetworkInterface()
	if activeInterface == "" {
		return "", errors.New("Failed getting active network interface")
	}

	if operation == "set" {
		setDns("185.51.200.2", "178.22.122.100")
		return "Success", nil
	} else {
		if err := resetDns(activeInterface); err != nil {
			return "", errors.New(fmt.Sprintf("Error resetting DNS for %s: %v", activeInterface, err))
		} else {
			return fmt.Sprintf("DNS settings reset successfully for %s\n", activeInterface), nil
		}
	}
}

func resetDns(interfaceName string) error {
	cmd := exec.Command("netsh", "interface", "ipv4", "set", "dnsservers", "name="+interfaceName, "source=dhcp")
	return cmd.Run()
}

func setDns(dns1 string, dns2 string) {
	var activeInterface = getActiveNetworkInterface()
	if activeInterface == "" {
		fmt.Print("Failed getting active network interface")
		return
	}

	commandText := fmt.Sprintf(`Set-DNSClientServerAddress -InterfaceAlias "%s" -ServerAddresses ("%s","%s")`,
		activeInterface, dns1, dns2)

	cmd := exec.Command("powershell", "-Command", commandText)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to execute command: %s, Output: %s\n", err, output)
		return
	}

	fmt.Println(output)

	fmt.Println("DNS settings updated successfully.")
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
