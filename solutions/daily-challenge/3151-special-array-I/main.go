package main

import "fmt"

func isArraySpecial(nums []int) bool {
	for i, x := range nums[1:] {
		if x%2 == nums[i]%2 {
			return false
		}
	}
	return true
}

func main() {
	testCases := []struct {
		nums []int
	}{
		{nums: []int{1}},
		{nums: []int{2, 1, 4}},
		{nums: []int{4, 3, 1, 6}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.nums)
		fmt.Printf("OUTPUT: %v\n\n", isArraySpecial(testCase.nums))
	}
}
