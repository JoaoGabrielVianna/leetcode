package main

import "fmt"

func countBadPairs(nums []int) (ans int64) {
	cnt := map[int]int{}
	for i, x := range nums {
		x = i - x
		ans += int64(i - cnt[x])
		cnt[x]++
	}
	return
}

func main() {
	testCases := []struct {
		nums []int
	}{
		{nums: []int{4, 1, 3, 3}},
		{nums: []int{1, 2, 3, 4, 5}},
	}
	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.nums)
		fmt.Printf("OUTPUT: %v\n\n", countBadPairs(testCase.nums))
	}
}
