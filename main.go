package main

import (
	"flag"
	"fmt"
	"net/netip"
	"snet/calc"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func printNetwork(prefix calc.Prefix) {
	na, _ := prefix.NetworkAddr()
	ba, _ := prefix.BroadcastAddr()
	mask, _ := prefix.Mask()
	hostCount, _ := prefix.HostsCount()

	t := table.New().
    BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("5"))).
	StyleFunc(func(row, col int) lipgloss.Style {
		switch {
			case row == 0:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("13")).Bold(true).PaddingLeft(1).PaddingRight(1)
			default: 
				return lipgloss.NewStyle().Foreground(lipgloss.Color("12")).PaddingLeft(1).PaddingRight(1)
		}
	}).Headers("Prefix", "Network", "Broadcast", "Mask", "Useable Hosts")

	t.Row(prefix.String(), na.String(), ba.String(), mask.String(), strconv.Itoa(hostCount))
	
	fmt.Println(t)
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
