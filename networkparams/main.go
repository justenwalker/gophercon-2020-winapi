package main

import (
	"fmt"
	"os"
)

// This implements the example program
// https://docs.microsoft.com/en-us/windows/win32/api/iptypes/ns-iptypes-fixed_info_w2ksp1#examples

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go syscall_windows.go

func main() {
	np, err := GetNetworkParams()
	if err != nil {
		fmt.Println("GetNetworkParams:", err)
		os.Exit(1)
	}
	fmt.Println("Host Name:", np.HostName)
	fmt.Println("Domain Name:", np.DomainName)
	fmt.Println("DNS Servers:")
	for _, addr := range np.DNSServerList {
		fmt.Printf("\t%s\n", addr)
	}
	fmt.Println("Node Type:", np.NodeType)
	if np.EnableRouting {
		fmt.Println("Routing: enabled")
	} else {
		fmt.Println("Routing: disabled")
	}
	if np.EnableProxy {
		fmt.Println("ARP Proxy: enabled")
	} else {
		fmt.Println("ARP Proxy: disabled")
	}
	if np.EnableDNS {
		fmt.Println("DNS: enabled")
	} else {
		fmt.Println("DNS: disabled")
	}
}
