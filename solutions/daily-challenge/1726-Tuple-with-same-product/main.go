package main

import "fmt"

func tupleSameProduct(nums []int) int {
	cnt := map[int]int{}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			x := nums[i] * nums[j]
			cnt[x]++
		}
	}
	ans := 0
	for _, v := range cnt {
		ans += v * (v - 1) / 2
	}
	return ans << 3
}

func main() {

	testCases := []struct {
		nums []int
	}{
		{nums: []int{2, 3, 4, 6}},
		{nums: []int{1, 2, 4, 5, 10}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.nums)
		fmt.Printf("OUTPUT: %v\n\n", tupleSameProduct(testCase.nums))
	}
}
