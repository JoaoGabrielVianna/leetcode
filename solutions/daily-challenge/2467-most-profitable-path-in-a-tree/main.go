package main

import "fmt"

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	ts := make([]int, n)
	for i := range ts {
		ts[i] = n
	}
	var dfs1 func(int, int, int) bool
	dfs1 = func(i, fa, t int) bool {
		if i == 0 {
			ts[i] = min(ts[i], t)
			return true
		}
		for _, j := range g[i] {
			if j != fa && dfs1(j, i, t+1) {
				ts[j] = min(ts[j], t+1)
				return true
			}
		}
		return false
	}
	dfs1(bob, -1, 0)
	ts[bob] = 0
	ans := -0x3f3f3f3f
	var dfs2 func(int, int, int, int)
	dfs2 = func(i, fa, t, v int) {
		if t == ts[i] {
			v += amount[i] >> 1
		} else if t < ts[i] {
			v += amount[i]
		}
		if len(g[i]) == 1 && g[i][0] == fa {
			ans = max(ans, v)
			return
		}
		for _, j := range g[i] {
			if j != fa {
				dfs2(j, i, t+1, v)
			}
		}
	}
	dfs2(0, -1, 0, 0)
	return ans
}

func main() {
	testCases := []struct {
		edges  [][]int
		bob    int
		amount []int
	}{
		{edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 6}}, bob: 0, amount: []int{1, 2, 3, 4, 5, 6, 7}},
		{edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 6}}, bob: 1, amount: []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.edges, testCase.bob, testCase.amount)
		fmt.Printf("OUTPUT: %v\n\n", mostProfitablePath(testCase.edges, testCase.bob, testCase.amount))
	}
}
