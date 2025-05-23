package main

import "fmt"

func minimizeXor(num1 int, num2 int) int {
	bits := 0
	for n := num2; n != 0; n >>= 1 {
		bits += n & 1
	}

	// Create bit arrays for num1 and res
	bitArray := make([]int, 32) // 32 bits to handle up to 32-bit integers
	for i := 0; i < 32; i++ {
		bitArray[i] = (num1 >> i) & 1
	}

	resArray := make([]int, 32)

	// Place `1`s in positions matching `num1` where possible
	for i := 31; i >= 0; i-- {
		if bits == 0 {
			break
		}
		if bitArray[i] == 1 {
			resArray[i] = 1
			bits--
		}
	}

	// Fill remaining bits with `1`s in lower positions
	for i := 0; i < 32; i++ {
		if bits == 0 {
			break
		}
		if resArray[i] == 0 {
			resArray[i] = 1
			bits--
		}
	}

	res := 0
	for i := 31; i >= 0; i-- {
		res = (res << 1) | resArray[i]
	}
	return res
}

func main() {

	testCases := []struct {
		num1 int
		num2 int
	}{
		{num1: 3, num2: 5},
		{num1: 1, num2: 12},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT NUM1", testCase.num1)
		fmt.Println("INPUT NUM2", testCase.num2)
		fmt.Println("OUTPUT", minimizeXor(testCase.num1, testCase.num2))
	}
}
