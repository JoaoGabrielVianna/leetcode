package main

import "fmt"

func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	idx := map[int][2]int{}
	for i := range mat {
		for j := range mat[i] {
			idx[mat[i][j]] = [2]int{i, j}
		}
	}
	row := make([]int, m)
	col := make([]int, n)
	for k := 0; ; k++ {
		x := idx[arr[k]]
		i, j := x[0], x[1]
		row[i]++
		col[j]++
		if row[i] == n || col[j] == m {
			return k
		}
	}
}

func main() {
	testCases := []struct {
		arr []int
		mat [][]int
	}{
		{arr: []int{1, 3, 4, 2}, mat: [][]int{{1, 4}, {2, 3}}},
		{arr: []int{2, 8, 7, 4, 1, 3, 5, 6, 9}, mat: [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT ARR:", testCase.arr)
		fmt.Println("INPUT MAT:", testCase.mat)
		fmt.Printf("OUTPUT: %d\n\n", firstCompleteIndex(testCase.arr, testCase.mat))

	}
}
