package main

import (
	"fmt"
	"sort"
)

func countDays(days int, meetings [][]int) (ans int) {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	last := 0
	for _, e := range meetings {
		st, ed := e[0], e[1]
		if last < st {
			ans += st - last - 1
		}
		last = max(last, ed)
	}
	ans += days - last
	return
}

func main() {
	testCases := []struct {
		days     int
		meetings [][]int
	}{
		{5, [][]int{{1, 2}, {3, 4}}},          // Espera-se 0 dias livres
		{10, [][]int{{1, 2}, {3, 5}, {6, 8}}}, // Espera-se 2 dias livres
		{7, [][]int{{1, 4}, {5, 6}}},          // Espera-se 1 dia livre
		{1, [][]int{{1, 1}}},                  // Espera-se 0 dias livres
	}

	for _, testCase := range testCases {
		fmt.Printf("INPUT: days = %d, meetings = %v\n", testCase.days, testCase.meetings)
		fmt.Printf("OUTPUT: %v\n\n", countDays(testCase.days, testCase.meetings))
	}
}
