package main

import "fmt"

func clearDigits(s string) string {
	stk := []byte{}
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, s[i])
		}
	}
	return string(stk)
}

func main() {
	testCases := []struct {
		s string
	}{
		{s: "abc"},
		{s: "cb34"},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.s)
		fmt.Printf("OUTPUT: %v\n\n", clearDigits(testCase.s))
	}
}
