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

func getHostRange(na netip.Addr, ba netip.Addr, hostCount int) string {
	if(!na.Is4()){
		return "Not Supported"
	}
	firstHost := na.Next().AsSlice()
	lastHost := ba.Prev().AsSlice()

	var hostRange string

	for i := 0; i < len(firstHost); i++ {
		if(firstHost[i] == lastHost[i]){
			hostRange += strconv.Itoa(int(firstHost[i]))
		} else if (firstHost[i] < lastHost[i]){
			hostRange += strconv.Itoa(int(firstHost[i])) + "-" + strconv.Itoa(int(lastHost[i]))
		} else {
			return "None"
		}

		if(i < len(firstHost) - 1) {
			hostRange += "."
		}
	}

	return hostRange + " (" + strconv.Itoa(hostCount) + ")"
}

func printNetworks(prefixes []calc.Prefix) {
	t := table.New().
    BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("5"))).
	StyleFunc(func(row, col int) lipgloss.Style {
		switch {
			case row == 0:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("13")).Bold(true).PaddingLeft(1).PaddingRight(1)
			case row > 1:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("12")).PaddingLeft(1).PaddingRight(1).PaddingTop(1)
			default: 
				return lipgloss.NewStyle().Foreground(lipgloss.Color("12")).PaddingLeft(1).PaddingRight(1)
		}
	}).Headers("Prefix", "Network", "Useable Hosts", "Broadcast", "Mask")

	for _, prefix := range prefixes {
		na, _ := prefix.NetworkAddr()
		ba, _ := prefix.BroadcastAddr()
		mask, _ := prefix.Mask()
		hc, _ := prefix.HostsCount()
		hostRange := getHostRange(na, ba, hc)

		t.Row(prefix.String(), na.String(), hostRange, ba.String(), mask.String())
	}
	
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
		prefixes := make([]calc.Prefix, 0, 1)
		prefixes = append(prefixes, calc.Prefix{Prefix: p})
		printNetworks(prefixes)
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
			prefixes := make([]calc.Prefix, 0, 1)
			prefixes = append(prefixes, calc.Prefix{Prefix: p})
			printNetworks(prefixes)
		} else {
			fmt.Println("Invalid network provided")
		}
	}
}
