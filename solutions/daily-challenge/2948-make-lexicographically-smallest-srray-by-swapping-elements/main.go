package main

import (
	"fmt"
	"slices"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return nums[i] - nums[j] })
	ans := make([]int, n)
	for i := 0; i < n; {
		j := i + 1
		for j < n && nums[idx[j]]-nums[idx[j-1]] <= limit {
			j++
		}
		t := slices.Clone(idx[i:j])
		slices.Sort(t)
		for k := i; k < j; k++ {
			ans[t[k-i]] = nums[idx[k]]
		}
		i = j
	}
	return ans
}

func main() {
	testCases := []struct {
		nums  []int
		limit int
	}{
		{nums: []int{1, 5, 3, 9, 8}, limit: 2},
		{nums: []int{1, 7, 6, 18, 2, 1}, limit: 3},
		{nums: []int{1, 7, 28, 19, 10}, limit: 3},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.nums, testCase.limit)
		fmt.Printf("OUTPUT: %v\n\n", lexicographicallySmallestArray(testCase.nums, testCase.limit))
	}
}
