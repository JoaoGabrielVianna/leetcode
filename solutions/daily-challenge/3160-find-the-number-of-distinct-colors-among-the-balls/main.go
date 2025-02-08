package main

import "fmt"

func queryResults(limit int, queries [][]int) (ans []int) {
	g := map[int]int{}
	cnt := map[int]int{}
	for _, q := range queries {
		x, y := q[0], q[1]
		cnt[y]++
		if v, ok := g[x]; ok {
			cnt[v]--
			if cnt[v] == 0 {
				delete(cnt, v)
			}
		}
		g[x] = y
		ans = append(ans, len(cnt))
	}
	return
}

func main() {

	testCases := []struct {
		limit   int
		queries [][]int
	}{
		{4, [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}},
		{4, [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}}},
	}

	for _, tc := range testCases {
		fmt.Println("INPUT:", tc.limit, tc.queries)
		fmt.Printf("OUTPUT: %v \n\n", queryResults(tc.limit, tc.queries))
	}

}
