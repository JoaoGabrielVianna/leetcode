package main

import "fmt"

func waysToSplitArray(nums []int) int {
	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0
	numberOfValidSplit := 0

	// Percorrer o array até o penúltimo elemento (último elemento não pode ser início de uma divisão válida)
	for i := 0; i < len(nums)-1; i++ {
		leftSum += nums[i]

		rightSum := totalSum - leftSum

		if leftSum >= rightSum {
			numberOfValidSplit++
		}
	}

	return numberOfValidSplit
}

func main() {
	num1 := []int{10, 4, -8, 7} // -> Matriz
	num2 := []int{2, 3, 1, 0}   // -> Matriz

	fmt.Println(waysToSplitArray(num1))
	fmt.Println(waysToSplitArray(num2))

	waysToSplitArray(num1)

}
