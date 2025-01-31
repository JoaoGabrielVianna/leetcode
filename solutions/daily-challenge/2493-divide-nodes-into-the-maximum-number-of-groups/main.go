package main

import "fmt"

func magnificentSets(n int, edges [][]int) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0]-1, e[1]-1
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	d := make([]int, n)
	for i := range d {
		q := []int{i}
		dist := make([]int, n)
		dist[i] = 1
		mx := 1
		root := i
		for len(q) > 0 {
			a := q[0]
			q = q[1:]
			root = min(root, a)
			for _, b := range g[a] {
				if dist[b] == 0 {
					dist[b] = dist[a] + 1
					mx = max(mx, dist[b])
					q = append(q, b)
				} else if abs(dist[b]-dist[a]) != 1 {
					return -1
				}
			}
		}
		d[root] = max(d[root], mx)
	}
	for _, x := range d {
		ans += x
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

	testCases := []struct {
		n     int
		edges [][]int
	}{
		{n: 6, edges: [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}},
		{n: 3, edges: [][]int{{1, 2}, {2, 3}, {3, 1}}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.n, testCase.edges)
		fmt.Printf("OUTPUT: %v\n\n", magnificentSets(testCase.n, testCase.edges))
	}
}
