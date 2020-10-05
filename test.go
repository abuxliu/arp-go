package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, _ := net.Interfaces()
	addresses, _ := net.InterfaceAddrs()
	ipNet := addresses[1].(*net.IPNet)
	device := interfaces[1].Name
	mac := interfaces[1].HardwareAddr
	ip := ipNet.IP.String()
	mask, _ := ipNet.Mask.Size()
	fmt.Printf("Device: %s, MAC: %s, IP: %s, MASK: %d.\n", device, mac, ip, mask)
}
