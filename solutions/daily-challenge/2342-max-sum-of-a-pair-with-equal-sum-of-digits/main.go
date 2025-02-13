package main

import "fmt"

func maximumSum(nums []int) int {
	d := [100]int{}
	ans := -1
	for _, v := range nums {
		x := 0
		for y := v; y > 0; y /= 10 {
			x += y % 10
		}
		if d[x] > 0 {
			ans = max(ans, d[x]+v)
		}
		d[x] = max(d[x], v)
	}
	return ans
}

func main() {
	testCases := []struct {
		nums []int
	}{
		{nums: []int{51, 71, 17, 42}},
		{nums: []int{11, 12, 13, 14}},
		{nums: []int{10, 12, 13, 14}},
	}
	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.nums)
		fmt.Printf("OUTPUT: %v\n\n", maximumSum(testCase.nums))
	}
}
