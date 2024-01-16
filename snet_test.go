package main

import "testing"

func TestBinaryToDecimal(t *testing.T){
	if binaryToDecimal("101") != 5 {
		t.Error("Expected 5, got ", binaryToDecimal("101"))
	}
	if binaryToDecimal("1") != 1 {
		t.Error("Expected 1, got ", binaryToDecimal("1"))
	}
	if binaryToDecimal("0") != 0 {
		t.Error("Expected 0, got ", binaryToDecimal("0"))
	}
	if binaryToDecimal("0000") != 0 {
		t.Error("Expected 0, got ", binaryToDecimal("0000"))
	}
	if binaryToDecimal("11111111") != 255 {
		t.Error("Expected 255, got ", binaryToDecimal("11111111"))
	}
	if binaryToDecimal("10110101") != 181 {
		t.Error("Expected 181, got ", binaryToDecimal("10110101"))
	}
}

func TestDecimalToBinary(t *testing.T){
	if decimalToBinary(5) != "101" {
		t.Error("Expected 101, got ", decimalToBinary(5))
	}
	if decimalToBinary(1) != "1" {
		t.Error("Expected 1, got ", decimalToBinary(1))
	}
	if decimalToBinary(0) != "0" {
		t.Error("Expected 0, got ", decimalToBinary(0))
	}
	if decimalToBinary(255) != "11111111" {
		t.Error("Expected 11111111, got ", decimalToBinary(255))
	}
	if decimalToBinary(181) != "10110101" {
		t.Error("Expected 10110101, got ", decimalToBinary(181))
	}
}

// t.Log("Similar to fmt.Println() and concurrently safe")
// t.Fail() // marks the test as failed but continues execution
// t.FailNow() // marks the test as failed and stops execution
// t.Error() // similar to t.Log() followed by t.Fail()
// t.Fatal() // similar to t.Log() followed by t.FailNow()
