package main

import (
	"flag"
	"fmt"
	"math"
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

// list of all hosts in the network (smaller mask bits take exponitially more time)
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

// count hosts in 0(1) time complexity
func (p Prefix) HostsCount() (int, error) {
	addr := p.Addr()
	bits := p.Bits()

	if (bits < 0 || bits > 128) {
		return 0, fmt.Errorf("invalid bit length")
	}

	var hostBits int
	if(addr.Is4()) {
		hostBits = 32 - bits
	} else if (addr.Is6()) {
		hostBits = 128 - bits
	} else {
		return 0, fmt.Errorf("invalid address")
	}

	hosts := math.Pow(2, float64(hostBits)) - 2

	return int(hosts), nil
}

func printNetwork(prefix Prefix) {
	na, _ := prefix.NetworkAddr()
	ba, _ := prefix.BroadcastAddr()
	hostCount, _ := prefix.HostsCount()

	fmt.Println("Network Address:", na)
	fmt.Println("Broadcast Address:", ba)
	fmt.Println("# of Hosts:", hostCount)
}

func addrToBits(addr netip.Addr) (int, error) {
	addrBytes := addr.AsSlice()

    bits := 0
    for _, b := range addrBytes {
        // Count the number of set bits in each byte
        for mask := byte(0x80); mask != 0; mask >>= 1 {
            if b&mask != 0 {
                bits++
            }
        }
    }

    return bits, nil
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
		prefix := Prefix{p}
		
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

		maskBits, _ := addrToBits(addr2)
		p := netip.PrefixFrom(addr1, maskBits)
		if(p.IsValid()) {
			prefix := Prefix{p}
			printNetwork(prefix)
		} else {
			fmt.Println("Invalid network provided")
		}
	}
}
