package main

import (
	"fmt"
	"strings"
)

func findDifferentBinaryString(nums []string) string {
	mask := 0
	for _, x := range nums {
		mask |= 1 << strings.Count(x, "1")
	}
	for i := 0; ; i++ {
		if mask>>i&1 == 0 {
			return strings.Repeat("1", i) + strings.Repeat("0", len(nums)-i)
		}
	}
}

func main() {
	testCases := []struct {
		nums []string
	}{
		{nums: []string{"01", "10", "11", "00"}},
		{nums: []string{"00", "01", "11", "10"}},
		{nums: []string{"111", "011", "001", "101", "100", "010", "110", "000"}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.nums)
		fmt.Printf("OUTPUT: %v\n\n", findDifferentBinaryString(testCase.nums))
	}
}
