package main

import (
	"flag"
	"fmt"
	"net/netip"
	"snet/calc"
)

func printNetwork(prefix calc.Prefix) {
	na, _ := prefix.NetworkAddr()
	ba, _ := prefix.BroadcastAddr()
	hostCount, _ := prefix.HostsCount()

	fmt.Println("Network Address:", na)
	fmt.Println("Broadcast Address:", ba)
	fmt.Println("# of Hosts:", hostCount)
}

func main() {
	flag.Parse()
	args := flag.Args()

	if(len(args) == 1) {
		p, err := netip.ParsePrefix(args[0])
		if err != nil{
			fmt.Println("Invalid network provided")
			return
		}
		prefix := calc.Prefix{p}
		
		printNetwork(prefix)
	} else if (len(args) > 1) {
		ip1 := args[0]
		ip2 := args[1]

		addr1, err := netip.ParseAddr(ip1)
		if err != nil {
			fmt.Println("Invalid address provided")
			return
		}

		addr2, err := netip.ParseAddr(ip2)
		if err != nil {
			fmt.Println("Invalid address provided")
			return
		}

		maskBits, _ := calc.AddrToBits(addr2)
		p := netip.PrefixFrom(addr1, maskBits)
		if(p.IsValid()) {
			prefix := calc.Prefix{p}
			printNetwork(prefix)
		} else {
			fmt.Println("Invalid network provided")
		}
	}
}
