package main

import "fmt"

func findThePrefixCommonArray(A []int, B []int) []int {
	res := make([]int, 0)
	a, b := 0, 0
	for i := 0; i < len(A); i++ {
		a |= 1 << A[i]
		b |= 1 << B[i]
		c := a & b
		res = append(res, bitCount(c))

	}

	return res
}

func bitCount(a int) int {
	res := 0
	for a != 0 {
		res += a & 1
		a >>= 1
	}
	return res
}

func main() {
	testCases := []struct {
		A []int
		B []int
	}{
		{A: []int{1, 3, 2, 4}, B: []int{3, 1, 2, 4}},
		{A: []int{2, 3, 1}, B: []int{3, 1, 2}},
	}

	for _, testCase := range testCases {
		fmt.Println("Input A:", testCase.A)
		fmt.Println("Input B:", testCase.B)
		fmt.Println("OUTPUT", findThePrefixCommonArray(testCase.A, testCase.B), "\n")
	}
}
