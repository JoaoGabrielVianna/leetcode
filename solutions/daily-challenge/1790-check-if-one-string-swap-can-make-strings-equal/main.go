package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {
	cnt := 0
	var c1, c2 byte
	for i := range s1 {
		a, b := s1[i], s2[i]
		if a != b {
			cnt++
			if cnt > 2 || (cnt == 2 && (a != c2 || b != c1)) {
				return false
			}
			c1, c2 = a, b
		}
	}
	return cnt != 1
}

func main() {
	testCases := []struct {
		s1 string
		s2 string
	}{
		{s1: "bank", s2: "kanb"},
		{s1: "attack", s2: "defend"},
		{s1: "kelb", s2: "kelb"},
		{s1: "abcd", s2: "dcba"},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.s1, testCase.s2)
		fmt.Printf("OUTPUT: %v\n\n", areAlmostEqual(testCase.s1, testCase.s2))
	}
}
