package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Look up the IP addresses for Google's public DNS hostname
	ips, err := net.LookupIP("dns.google")
	if err != nil {
		fmt.Println("Error looking up IP:", err)
		return
	}

	// Format the output string
	output := fmt.Sprintf("Google DNS IP addresses:\n")
	for _, ip := range ips {
		output += fmt.Sprintf("- %s\n", ip.String())
	}

	// Write the result to a text file named google_dns.txt
	err = os.WriteFile("google_dns.txt", []byte(output), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Successfully saved Google DNS IP to google_dns.txt!")
}
