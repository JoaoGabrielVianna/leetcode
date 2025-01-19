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
