package main

import (
	"fmt"
	"math"
)

func minOperations(boxes string) []int {
	n := len(boxes)
	answer := make([]int, n)

	for i := 0; i < n; i++ {
		totalOperatiopns := 0
		for j := 0; j < n; j++ {
			if boxes[j] == '1' {
				totalOperatiopns += int(math.Abs(float64(i - j)))
			}
		}
		answer[i] = totalOperatiopns
	}

	return answer
}

func main() {
	boxes := "110"
	result := minOperations(boxes)
	fmt.Println(result) // Output: [1, 1, 3]

	boxes2 := "001011"
	result2 := minOperations(boxes2)
	fmt.Println(result2) // Output: [11, 8, 5, 4, 3, 4]
}

// 1 1 0
