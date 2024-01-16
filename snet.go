package main

import (
	"fmt"
	"regexp"
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

func main() {
	input := "192.168.2.3"

	cidrRegex := regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3}\/\d{1,2}`)
	ipRegex := regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3}`)

	if cidrRegex.MatchString(input) {
		fmt.Println("ipv4 address with subnet mask")
	} else if ipRegex.MatchString(input) {
		fmt.Println("ipv4 address")
	} else {
		fmt.Println("invalid input")
	}

}
