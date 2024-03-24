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

func getNetworkAddrBytes(ipBytes []byte, maskBytes []byte) []byte {
	naBytes := make([]byte, len(ipBytes))
    for i := range ipBytes {
		// Bitwise AND
		naBytes[i] = ipBytes[i] & maskBytes[i]
    }
	return naBytes
}

func getBroadcastAddrBytes(ipBytes []byte, maskBytes []byte) []byte {
	baBytes := make([]byte, len(ipBytes))
    for i := range ipBytes {
		// Bitwise OR
		baBytes[i] = ipBytes[i] | ^maskBytes[i]
    }
	return baBytes
}

// network stuct based on netip.Prefix to add custom methods
type Prefix struct {
	netip.Prefix
	// addr
	// bits
}

// network mask of the network
func (p Prefix) Mask() (netip.Addr, error) {
	addr := p.Addr()
	bits := p.Bits()

	mask, err := bitsToMask(bits, addr.Is4())
	if err != nil {
		return netip.IPv4Unspecified(), fmt.Errorf("invalid network mask")
	}

	return mask, nil
}

// network address of the network ex. 192.168.20.15/23 -> 192.168.20.0
func (p Prefix) NetworkAddr() (netip.Addr, error) {
	addr := p.Addr()
	mask, err := p.Mask()
	if err != nil {
		return netip.IPv4Unspecified(), err
	}
	
	naBytes := getNetworkAddrBytes(addr.AsSlice(), mask.AsSlice())

	na, ok := netip.AddrFromSlice(naBytes)
	if !ok {
		return netip.IPv4Unspecified(), fmt.Errorf("invalid network address")
	}

	return na, nil
}

// broadcast address of the network ex. 192.168.20.15/23 -> 192.168.21.255
func (p Prefix) BroadcastAddr() (netip.Addr, error) {
	addr := p.Addr()
	mask, err := p.Mask()
	if err != nil {
		return netip.IPv4Unspecified(), err
	}
	
	baBytes := getBroadcastAddrBytes(addr.AsSlice(), mask.AsSlice())

	ba, ok := netip.AddrFromSlice(baBytes)
	if !ok {
		return netip.IPv4Unspecified(), fmt.Errorf("invalid network address")
	}

	return ba, nil
}

// list of all hosts in the network
func (p Prefix) Hosts() ([]netip.Addr, error) {
	addr := p.Addr()
	mask, err := p.Mask()
	if err != nil {
		return []netip.Addr{}, err
	}

	naBytes := getNetworkAddrBytes(addr.AsSlice(), mask.AsSlice())
	baBytes := getBroadcastAddrBytes(addr.AsSlice(), mask.AsSlice())

	na, _ := netip.AddrFromSlice(naBytes)
	ba, _ := netip.AddrFromSlice(baBytes)

	cursor := na
	hosts := []netip.Addr{}
	for cursor.Less(ba.Prev()) {
		hosts = append(hosts, cursor.Next())
		cursor = cursor.Next()
	}

	return hosts, nil
}


func main() {
	addr,_ := netip.ParseAddr("172.12.20.15")

	prefix := Prefix{netip.PrefixFrom(addr, 23)}
	na, _ := prefix.NetworkAddr()
	ba, _ := prefix.BroadcastAddr()
	hosts, _ := prefix.Hosts()

	var netType string
	if(addr.IsPrivate()) {
		netType = "Private"
	} else {
		netType = "Public"
	}

	fmt.Println("Network Address:", na)
	fmt.Println("Broadcast Address:", ba)
	fmt.Println("Hosts:", len(hosts))
	fmt.Println("Type:", netType)
}
