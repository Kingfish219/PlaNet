package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// interfaceName := "Local Area Connection"

	// dns1 := "8.8.8.8"
	dns2 := "8.8.4.4"

	// commandText := fmt.Sprintf(`Set-DNSClientServerAddress -InterfaceAlias "%s" -ServerAddresses ("%s","%s")`, interfaceName, dns1, dns2)
	commandText := fmt.Sprintf(`ping "%s"`, dns2)

	cmd := exec.Command("powershell", "-Command", commandText)

	// Execute the PowerShell command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to execute command: %s, Output: %s\n", err, output)
		return
	}

	fmt.Println(output)
	fmt.Println("DNS settings updated successfully.")
}
