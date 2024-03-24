package main

import (
	"fmt"
	"net"
	"net/netip"
)

func bitsToMask(bits int, ipv4 bool) (netip.Addr, error) {
	if bits < 0 || bits > 128 {
		return netip.IPv4Unspecified(), fmt.Errorf("invalid bit length")
	}

	var ip net.IP
	var addr netip.Addr

	if(bits <= 32 && ipv4 == true) {
		mask := net.CIDRMask(bits, 32)
		ip = net.IP(mask).To4()
	} else {
		mask := net.CIDRMask(bits, 128)
		ip = net.IP(mask).To16()
	}
	addr, _ = netip.ParseAddr(ip.String())
	return addr, nil
}

func getNetDetails(ip netip.Addr, bits int) (netip.Addr, netip.Addr, error) {
	prefix := netip.PrefixFrom(ip, bits)
	if prefix.IsValid() == false {
		return netip.IPv4Unspecified(), netip.IPv4Unspecified(), fmt.Errorf("invalid network")
	}

	mask, _ := bitsToMask(bits, ip.Is4())

	ipBytes := ip.AsSlice()
	maskBytes := mask.AsSlice()

	// Network Address (Bitwise AND)
	naBytes := make([]byte, len(ipBytes))
    for i := range ipBytes {
		naBytes[i] = ipBytes[i] & maskBytes[i]
    }

	// Broadcast Address (Bitwise OR)
	baBytes := make([]byte, len(ipBytes))
	for i := range ipBytes {
		baBytes[i] = ipBytes[i] | ^maskBytes[i]
    }

	na, _ := netip.AddrFromSlice(naBytes)
	ba, _ := netip.AddrFromSlice(baBytes)
	
	return na, ba, nil
}


func main() {
	addr,err := netip.ParseAddr("192.168.20.12")

	if err != nil {
		fmt.Println("Error:", err)
	}

	na, ba, err := getNetDetails(addr, 23)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Network Address:", na)
	fmt.Println("Broadcast Address:", ba)
}
