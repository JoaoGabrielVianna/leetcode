package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if p[x] != x {
			p[x] = find(p[x])
		}
		return p[x]
	}
	for i := 0; ; i++ {
		pa, pb := find(edges[i][0]-1), find(edges[i][1]-1)
		if pa == pb {
			return edges[i]
		}
		p[pa] = pb
	}
}

func main() {
	testCases := []struct {
		edges [][]int
	}{
		{edges: [][]int{{1, 2}, {1, 3}, {2, 3}}},
		{edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.edges)
		fmt.Printf("OUTPUT: %v\n\n", findRedundantConnection(testCase.edges))
	}
}
