package main

import (
	"fmt"
	"net"
)


func binaryToDecimal(binary string) int {
	decimal := 0
	for i := 0; i < len(binary); i++ {
		decimal = decimal * 2
		if binary[i] == '1' {
			decimal = decimal + 1
		}
	}
	return decimal
}

func decimalToBinary(decimal int) string {
	if decimal == 0 { return "0" }
	binary := ""
	for decimal > 0 {
		remainder := decimal % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		decimal = decimal / 2
	}
	return binary
}


func getIPType(address string) string {
	ip := net.ParseIP(address)
	if ip == nil { return "invalid" }

	if ip.To4() != nil {
		return "ipv4"
	} else if ip.To16() != nil {
		return "ipv6"
	}
	return "unknown"
}

func main() {
	input := "192.168.2.3"

	fmt.Println(getIPType(input))


}
