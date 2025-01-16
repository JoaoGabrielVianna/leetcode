package main

import "fmt"

func xorAllNums(nums1 []int, nums2 []int) int {
	N1, N2 := len(nums1), len(nums2)

	res := 0

	if N1%2 == 1 {
		for _, v := range nums2 {
			res ^= v
		}
	}
	if N2%2 == 1 {
		for _, v := range nums1 {
			res ^= v
		}
	}
	return res
}

func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
	}{
		{nums1: []int{2, 1, 3}, nums2: []int{10, 2, 5, 0}},
		{nums1: []int{1, 2}, nums2: []int{3, 4}},
	}

	for _, tc := range testCases {

		fmt.Println("INPUT num1:", tc.nums1)
		fmt.Println("INPUT num2:", tc.nums2)
		fmt.Println("OUTPUT:", xorAllNums(tc.nums1, tc.nums2), "\n")
	}

}
