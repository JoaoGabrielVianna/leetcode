package main

import "fmt"

func findMaxFish(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	dirs := [5]int{-1, 0, 1, 0, -1}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		cnt := grid[i][j]
		grid[i][j] = 0
		for k := 0; k < 4; k++ {
			x, y := i+dirs[k], j+dirs[k+1]
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] > 0 {
				cnt += dfs(x, y)
			}
		}
		return cnt
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return
}

func main() {
	testCases := []struct {
		grid [][]int
	}{
		{grid: [][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}},
		{grid: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}}},
	}
	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.grid)
		fmt.Printf("OUTPUT: %v\n\n", findMaxFish(testCase.grid))
	}
}
