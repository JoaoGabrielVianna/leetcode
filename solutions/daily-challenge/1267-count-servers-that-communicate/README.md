# 1267. Count Servers that Communicate

**`Medium`**

You are given a map of a server center, represented as a `m * n` integer matrix `grid`, where 1 means that on that cell there is a server and 0 means that it is no server. Two servers are said to communicate if they are on the same row or on the same column.

Return the number of servers that communicate with any other server.

### Example 1:
![Grid 1](assets/untitled-diagram-6.jpg)

**Input:** grid = [[1,0],[0,1]]

**Output:** 0

**Explanation:** 

No servers can communicate with others.

### Example 2:
![Grid 2](assets/untitled-diagram-4.jpg)

**Input:** grid = [[1,0],[1,1]]

**Output:** 3

**Explanation:** 

All three servers can communicate with at least one other server.

### Example 3:
![Grid 3](assets/untitled-diagram-1-3.jpg)

**Input:** grid = [[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]

**Output:** 4

**Explanation:** 

The two servers in the first row can communicate with each other. The two servers in the third column can communicate with each other. The server at right bottom corner can't communicate with any other server.
 

### Constraints:

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m <= 250`
- `1 <= n <= 250`
- `grid[i][j] == 0 or 1`

### Solution

```go
func countServers(grid [][]int) int {
	ROWS, COLS := len(grid), len(grid[0])
	rowCounts := make([]int, ROWS)
	colCounts := make([]int, COLS)
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if grid[r][c] == 1 {
				rowCounts[r]++
				colCounts[c]++
			}
		}
	}
	var res int = 0
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if grid[r][c] == 1 && (rowCounts[r]+colCounts[c] > 2) {
				res++
			}
		}
	}
	return res
}
```

### Test

```go
func main() {

	testCases := []struct {
		grid [][]int
	}{
		{grid: [][]int{{1, 0}, {0, 1}}},
		{grid: [][]int{{1, 0}, {1, 1}}},
		{grid: [][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.grid)
		fmt.Printf("OUTPUT: %v\n\n", countServers(testCase.grid))
	}
}
```