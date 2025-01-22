# 1765. Map of Highest Peak

**`Medium`**

You are given an integer matrix isWater of size m x n that represents a map of land and water cells.

- If isWater[i][j] == 0, cell (i, j) is a land cell.
- If isWater[i][j] == 1, cell (i, j) is a water cell.
You must assign each cell a height in a way that follows these rules:

- The height of each cell must be non-negative.
- If the cell is a water cell, its height must be 0.
- Any two adjacent cells must have an absolute height difference of at most 1. A cell is adjacent to another cell if the former is directly north, east, south, or west of the latter (i.e., their sides are touching).
Find an assignment of heights such that the maximum height in the matrix is maximized.

Return an integer matrix height of size m x n where height[i][j] is cell (i, j)'s height. If there are multiple solutions, return any of them.

### Example 1:
![example 1](assets/screenshot-2021-01-11-at-82045-am.png)

**Input:** isWater = [[0,1],[0,0]]

**Output:** [[1,0],[2,1]]

**Explanation:** 

The image shows the assigned heights of each cell.
The blue cell is the water cell, and the green cells are the land cells.

### Example 2:
![example 2](assets/screenshot-2021-01-11-at-82050-am.png)

**Input:** isWater = [[0,0,1],[1,0,0],[0,0,0]]

**Output:** [[1,1,0],[0,1,1],[1,2,2]]

**Explanation:** 

A height of 2 is the maximum possible height of any assignment.
Any height assignment that has a maximum height of 2 while still meeting the rules will also be accepted.
 

### Constraints:

- `m == isWater.length`
- `n == isWater[i].length`
- `1 <= m, n <= 1000`
- `isWater[i][j]` is `0` or `1`.
- There is at least one water cell.

### Solutions

```go
func highestPeak(isWater [][]int) [][]int {
	ROWS, COLS := len(isWater), len(isWater[0])
	directions := [][]int{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}
	res := make([][]int, ROWS)
	for i := 0; i < ROWS; i++ {
		res[i] = make([]int, COLS)
	}
	visited := make(map[[2]int]bool)
	bfs := list.New()
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if isWater[r][c] == 1 {
				bfs.PushBack([3]int{0, r, c})
				visited[[2]int{r, c}] = true
			}
		}
	}

	for bfs.Len() > 0 {
		elem := bfs.Front()
		bfs.Remove(elem)
		current := elem.Value.([3]int)
		num, x, y := current[0], current[1], current[2]

		res[x][y] = num

		for _, dir := range directions {
			nr, nc := x+dir[0], y+dir[1]
			if nr < 0 || nr >= ROWS || nc < 0 || nc >= COLS || visited[[2]int{nr, nc}] {
				continue
			}
			visited[[2]int{nr, nc}] = true
			bfs.PushBack([3]int{num + 1, nr, nc})
		}
	}

	return res
}
```

### Test

```go
func main() {
	testCases := []struct {
		isWater [][]int
	}{
		{isWater: [][]int{{0, 1}, {0, 0}}},
		{isWater: [][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.isWater)
		fmt.Printf("OUTPUT: %d\n\n", highestPeak(testCase.isWater))

	}
}
```