package main

import "fmt"

func longestMonotonicSubarray(nums []int) int {
	ans := 1
	t := 1
	for i, x := range nums[1:] {
		if nums[i] < x {
			t++
			ans = max(ans, t)
		} else {
			t = 1
		}
	}
	t = 1
	for i, x := range nums[1:] {
		if nums[i] > x {
			t++
			ans = max(ans, t)
		} else {
			t = 1
		}
	}
	return ans
}

func main() {

	testCases := []struct {
		nums []int
	}{
		{nums: []int{1, 4, 3, 3, 2}},
		{nums: []int{3, 3, 3, 3}},
		{nums: []int{3, 2, 1}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.nums)
		fmt.Printf("OUTPUT: %v\n\n", longestMonotonicSubarray(testCase.nums))
	}
}
