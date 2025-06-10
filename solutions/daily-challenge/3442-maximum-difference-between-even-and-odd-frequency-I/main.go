package main

import "fmt"

func maxDifference(s string) int {
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}
	oddFreqs := []int{}
	evenFreqs := []int{}

	for _, count := range freq {
		if count%2 == 1 {
			oddFreqs = append(oddFreqs, count)
		} else {
			evenFreqs = append(evenFreqs, count)
		}
	}

	maxOdd := max(oddFreqs)
	minEven := min(evenFreqs)

	return maxOdd - minEven

}

func max(slice []int) int {
	m := slice[0]
	for _, v := range slice {
		if v > m {
			m = v
		}
	}
	return m
}

func min(slice []int) int {
	m := slice[0]
	for _, v := range slice {
		if v < m {
			m = v
		}
	}
	return m
}

func main() {
	testCases := []struct {
		s string
	}{
		{s: "aaaaabbc"},
		{s: "abcabcab"},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT: ", testCase.s)
		fmt.Println("OUTPUT: ", maxDifference(testCase.s))
	}
}
