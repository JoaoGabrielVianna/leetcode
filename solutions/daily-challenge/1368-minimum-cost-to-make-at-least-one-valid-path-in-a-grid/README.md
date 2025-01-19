# 1368. Minimum Cost to Make at Least One Valid Path in a Grid

**`Hard`**

Given an `m x n` grid. Each cell of the grid has a sign pointing to the next cell you should visit if you are currently in this cell. The sign of `grid[i][j]` can be:

- `1` which means go to the cell to the right. (i.e go from `grid[i][j]` to `grid[i][j + 1]`)
- `2` which means go to the cell to the left. (i.e go from `grid[i][j]` to `grid[i][j - 1]`)
- `3` which means go to the lower cell. (i.e go from `grid[i][j]` to `grid[i + 1][j]`)
- `4` which means go to the upper cell. (i.e go from `grid[i][j]` to `grid[i - 1][j]`)
Notice that there could be some signs on the cells of the grid that point outside the grid.

You will initially start at the upper left cell `(0, 0)`. A valid path in the grid is a path that starts from the upper left cell `(0, 0)` and ends at the bottom-right cell `(m - 1, n - 1)` following the signs on the grid. The valid path does not have to be the shortest.

You can modify the sign on a cell with `cost = 1`. You can modify the sign on a cell one time only.

Return the minimum cost to make the grid have at least one valid path.

 

### Example 1:
![Grid 1](/assets/grid1.png)

**Input:** grid = [[1,1,1,1],[2,2,2,2],[1,1,1,1],[2,2,2,2]]

**Output:** 3

**Explanation:** 

You will start at point (0, 0).
The path to (3, 3) is as follows. (0, 0) --> (0, 1) --> (0, 2) --> (0, 3) change the arrow to down with cost = 1 --> (1, 3) --> (1, 2) --> (1, 1) --> (1, 0) change the arrow to down with cost = 1 --> (2, 0) --> (2, 1) --> (2, 2) --> (2, 3) change the arrow to down with cost = 1 --> (3, 3)
The total cost = 3.

### Example 2:
![Grid 1](/assets/grid2.png)

**Input:** grid = [[1,1,3],[3,2,2],[1,1,4]]

**Output:** 0

**Explanation:** 

You can follow the path from (0, 0) to (2, 2).

### Example 3:
![Grid 1](/assets/grid3.png)

**Input:** grid = [[1,2],[4,3]]
**Output:** 1
 

**Constraints:**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 100`
- `1 <= grid[i][j] <= 4`

### Solution

```go
package main

import (
	"container/list"
	"fmt"
)

func minCost(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	deque := list.New()
	deque.PushFront([3]int{0, 0, 0})
	visited := make(map[[2]int]bool)

	for deque.Len() > 0 {
		curr := deque.Remove(deque.Front()).([3]int)
		cost, x, y := curr[0], curr[1], curr[2]

		if x == rows-1 && y == cols-1 {
			return cost
		}

		if visited[[2]int{x, y}] {
			continue
		}
		visited[[2]int{x, y}] = true

		for i, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]

			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && !visited[[2]int{nx, ny}] {
				if i == grid[x][y]-1 {
					deque.PushFront([3]int{cost, nx, ny})
				} else {
					deque.PushBack([3]int{cost + 1, nx, ny})
				}
			}
		}
	}
	return -1
}
```

### Test

```go
func main() {
	testCases := []struct {
		grid [][]int
	}{
		{grid: [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}},
		{grid: [][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}},
		{grid: [][]int{{1, 2}, {4, 3}}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.grid)
		fmt.Println("OUTPUT:", minCost(testCase.grid), "\n")
	}
}

```