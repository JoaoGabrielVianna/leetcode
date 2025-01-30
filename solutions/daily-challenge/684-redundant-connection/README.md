# 684. Redundant Connection

**`Medium`**

In this problem, a tree is an undirected graph that is connected and has no cycles.

You are given a graph that started as a tree with n nodes labeled from 1 to n, with one additional edge added. The added edge has two different vertices chosen from 1 to n, and was not an edge that already existed. The graph is represented as an array edges of length n where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the graph.

Return an edge that can be removed so that the resulting graph is a tree of n nodes. If there are multiple answers, return the answer that occurs last in the input.

### Example 1:
![Example 1](assets/reduntant1-1-graph.jpg)

**Input:** edges = [[1,2],[1,3],[2,3]]

**Output:** [2,3]

### Example 2:
![Example 1](assets/reduntant1-2-graph.jpg)

**Input:** edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]

**Output:** [1,4]

### Constraints:

- `n == edges.length`
- `3 <= n <= 1000`
- `edges[i].length == 2`
- `1 <= ai < bi <= edges.length`
- `ai != bi`
- There are no repeated edges.
- The given graph is connected.

### Solution

```go
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
```

### Test

```go
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

```