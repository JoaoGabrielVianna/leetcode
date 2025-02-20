package main

import (
	"fmt"
	"sort"
)

func getHappyString(n int, k int) string {
	var result []string
	var backtrack func(curr string)

	backtrack = func(curr string) {
		if len(curr) == n {
			result = append(result, curr)
			return
		}
		for _, ch := range "abc" {
			if len(curr) == 0 || curr[len(curr)-1] != byte(ch) {
				backtrack(curr + string(ch))
			}
		}
	}

	backtrack("")
	sort.Strings(result)

	if k > len(result) {
		return ""
	}
	return result[k-1]
}

func main() {
	testCases := []struct {
		n int
		k int
	}{
		{n: 1, k: 3},
		{n: 2, k: 2},
		{n: 3, k: 3},
	}

	for i, tc := range testCases {
		fmt.Println("INPUT:", i+1, tc.n, tc.k)
		fmt.Printf("OUTPUT: %v\n\n", getHappyString(tc.n, tc.k))
	}
}
