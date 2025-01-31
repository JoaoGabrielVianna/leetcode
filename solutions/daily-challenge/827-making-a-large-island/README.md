# 827. Making A Large Island

**`Hard`**

You are given an n x n binary matrix grid. You are allowed to change at most one 0 to be 1.

Return the size of the largest island in grid after applying this operation.

An island is a 4-directionally connected group of 1s.

### Example 1:

**Input:** grid = [[1,0],[0,1]]

**Output:** 3

**Explanation:** 

Change one 0 to 1 and connect two 1s, then we get an island with area = 3.

### Example 2:

**Input:** grid = [[1,1],[1,0]]

**Output:** 4

**Explanation:** 

Change the 0 to 1 and make the island bigger, only one island with area = 4.

### Example 3:

**Input:** grid = [[1,1],[1,1]]

**Output:** 4

**Explanation:** 

Can't change any 0 to 1, only one island with area = 4.
 

### Constraints:

- `n == grid.length`
- `n == grid[i].length`
- `1 <= n <= 500`
- `grid[i][j] is either 0 or 1.`

### Solution

```go
func largestIsland(grid [][]int) int {
	n := len(grid)
	p := make([][]int, n)
	for i := range p {
		p[i] = make([]int, n)
	}
	cnt := make([]int, n*n+1)
	dirs := []int{-1, 0, 1, 0, -1}
	root := 0
	ans := 0

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		p[i][j] = root
		cnt[root]++
		for k := 0; k < 4; k++ {
			x := i + dirs[k]
			y := j + dirs[k+1]
			if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == 1 && p[x][y] == 0 {
				dfs(x, y)
			}
		}
		return cnt[root]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && p[i][j] == 0 {
				root++
				ans = max(ans, dfs(i, j))
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				s := make(map[int]struct{})
				for k := 0; k < 4; k++ {
					x := i + dirs[k]
					y := j + dirs[k+1]
					if x >= 0 && x < n && y >= 0 && y < n {
						s[p[x][y]] = struct{}{}
					}
				}
				t := 1
				for x := range s {
					t += cnt[x]
				}
				ans = max(ans, t)
			}
		}
	}
	return ans
}
```


### Test

```go
func main() {

	testCases := []struct {
		grid [][]int
	}{
		{
			grid: [][]int{
				{1, 0},
				{0, 1},
			},
		},
		{
			grid: [][]int{
				{1, 1},
				{1, 0},
			},
		},
		{
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
		},
	}

	for _, tc := range testCases {

		fmt.Println("INPUT:", tc.grid)
		fmt.Printf("OUTPUT: %v\n\n", largestIsland(tc.grid))

	}
}
```