package main

import (
	"fmt"
	"net"
)

func main() {
	host := "www.example.com"

	// Parse the host
	_, _, err := net.ParseCIDR(host)
	if err == nil {
		// Valid host format
		fmt.Println("Valid host/domain")
	} else {
		// Invalid format
		fmt.Println("Invalid host/domain")
	}
	lookupHost, err := net.LookupHost(host)
	if err == nil {
		println("ture")
	}
	fmt.Printf("%s\n", lookupHost)
	println(isValidHost("8.8.8.8"))            // true
	println(isValidHost("example.com"))        // true
	println(isValidHost("http://example.com")) // true
	println(isValidHost("invalidhost"))        // false
	println(isValidHost("example.com.ico"))
	println(isValidHost("example.com/login"))
}

func isValidHost(host string) bool {
	// Parse as IP address
	if ip := net.ParseIP(host); ip != nil {
		return true
	}

	// Parse as hostname
	if _, err := net.LookupHost(host); err == nil {
		return true
	}

	// Parse as URL
	//if _, err := url.Parse(host); err == nil {
	//	return true
	//}

	return false
}
