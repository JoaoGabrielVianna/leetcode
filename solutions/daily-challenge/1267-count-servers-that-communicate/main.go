package main

import (
	"fmt"
)

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
