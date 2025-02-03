package main

import "fmt"

func check(nums []int) bool {
	cnt := 0
	for i, v := range nums {
		if v > nums[(i+1)%len(nums)] {
			cnt++
		}
	}
	return cnt <= 1
}

func main() {

	testCases := []struct {
		nums []int
	}{
		{nums: []int{3, 4, 5, 1, 2}},
		{nums: []int{2, 1, 3, 4}},
		{nums: []int{1, 2, 3}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.nums)
		fmt.Printf("OUTPUT: %v\n\n", check(testCase.nums))
	}
}
