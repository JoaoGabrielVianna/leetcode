package main

import (
	"fmt"
	"sort"
)

func putMarbles(weights []int, k int) (ans int64) {
	n := len(weights)
	arr := make([]int, n-1)
	for i, w := range weights[:n-1] {
		arr[i] = w + weights[i+1]
	}
	sort.Ints(arr)
	for i := 0; i < k-1; i++ {
		ans += int64(arr[n-2-i] - arr[i])
	}
	return
}

func main() {
	testCases := []struct {
		weights []int
		k       int
		output  int64
	}{
		{weights: []int{1, 3, 5, 1}, k: 2, output: 4},
		{weights: []int{1, 3}, k: 1, output: 0},
		{weights: []int{10, 20, 30, 40, 50}, k: 3, output: 60},
		{weights: []int{5, 5, 5, 5}, k: 2, output: 0},
		{weights: []int{1, 2, 3, 4, 5}, k: 4, output: 6},
	}
	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.weights, testCase.k)
		fmt.Printf("OUTPUT: %v\n", putMarbles(testCase.weights, testCase.k))
		fmt.Printf("EXPECTED: %v\n\n", testCase.output)
	}
}
