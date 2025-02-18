package main

import "fmt"

func numTilePossibilities(tiles string) int {
	cnt := [26]int{}
	for _, c := range tiles {
		cnt[c-'A']++
	}
	var dfs func(cnt [26]int) int
	dfs = func(cnt [26]int) (res int) {
		for i, x := range cnt {
			if x > 0 {
				res++
				cnt[i]--
				res += dfs(cnt)
				cnt[i]++
			}
		}
		return
	}
	return dfs(cnt)
}

func main() {
	testCases := []struct {
		tiles string
	}{
		{tiles: "AAB"},
		{tiles: "AAABBC"},
	}
	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.tiles)
		fmt.Printf("OUTPUT: %v\n\n", numTilePossibilities(testCase.tiles))
	}
}
