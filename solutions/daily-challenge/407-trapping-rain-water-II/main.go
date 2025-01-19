package main

import (
	"container/heap"
	"fmt"
)

func trapRainWater(hs [][]int) int {
	if len(hs) < 3 || len(hs[0]) < 3 {
		return 0
	}
	m, n := len(hs), len(hs[0])

	pq := make(priorityQueue, 0, m*2+n*2)

	isVisited := make([][]bool, m)
	for i := range isVisited {
		isVisited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		isVisited[i][0] = true
		isVisited[i][n-1] = true
		pq = append(pq, cell{row: i, col: 0, height: hs[i][0]})
		pq = append(pq, cell{row: i, col: n - 1, height: hs[i][n-1]})
	}
	for j := 0; j < n; j++ {
		isVisited[0][j] = true
		isVisited[m-1][j] = true
		pq = append(pq, cell{row: 0, col: j, height: hs[0][j]})
		pq = append(pq, cell{row: m - 1, col: j, height: hs[m-1][j]})
	}

	heap.Init(&pq)

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	vol := 0

	for len(pq) > 0 {

		c := heap.Pop(&pq).(cell)

		for _, d := range dirs {
			i := c.row + d[0]
			j := c.col + d[1]

			if 0 <= i && i < m && 0 <= j && j < n && !isVisited[i][j] {
				isVisited[i][j] = true

				vol += max(0, c.height-hs[i][j])

				heap.Push(&pq, cell{row: i, col: j, height: max(hs[i][j], c.height)})
			}
		}
	}

	return vol
}

type cell struct {
	row, col, height int
}

type priorityQueue []cell

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].height < pq[j].height
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(cell)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	testCases := []struct {
		hs [][]int
	}{
		{hs: [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}},
		{hs: [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.hs)
		fmt.Printf("OUTPUT: %d\n\n", trapRainWater(testCase.hs))

	}
}
