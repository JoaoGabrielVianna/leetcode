package main

import "fmt"

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) (ans []bool) {
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
	}
	for _, p := range prerequisites {
		f[p[0]][p[1]] = true
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[i][j] = f[i][j] || (f[i][k] && f[k][j])
			}
		}
	}
	for _, q := range queries {
		ans = append(ans, f[q[0]][q[1]])
	}
	return
}

func main() {
	testCases := []struct {
		n             int
		prerequisites [][]int
		queries       [][]int
		ans           []bool
	}{
		{2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}, []bool{false, true}},
		{2, [][]int{}, [][]int{{1, 0}, {0, 1}}, []bool{false, false}},
		{3, [][]int{{1, 2}, {1, 0}, {2, 0}}, [][]int{{1, 0}, {1, 2}}, []bool{false, true}},
	}
	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.n)
		fmt.Println("INPUT:", testCase.prerequisites)
		fmt.Println("INPUT:", testCase.queries)
		fmt.Printf("OUTPUT: %v\n\n", checkIfPrerequisite(testCase.n, testCase.prerequisites, testCase.queries))
	}
}
