package main

import "fmt"

func doesValidArrayExist(derived []int) bool {
	res := 0
	for _, num := range derived {
		res ^= num
	}
	return res == 0
}
func main() {
	// EU QUERO UM STRUCT DE testCases
	testCases := []struct {
		input []int
	}{
		{[]int{1, 2, 3, 1}},
		{[]int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
	}
	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.input)
		fmt.Println("OUTPUT:", doesValidArrayExist(testCase.input))
	}
}
