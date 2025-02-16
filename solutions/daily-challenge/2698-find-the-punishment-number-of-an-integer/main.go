package main

import (
	"fmt"
	"strconv"
)

func punishmentNumber(n int) (ans int) {
	var check func(string, int, int) bool
	check = func(s string, i, x int) bool {
		m := len(s)
		if i >= m {
			return x == 0
		}
		y := 0
		for j := i; j < m; j++ {
			y = y*10 + int(s[j]-'0')
			if y > x {
				break
			}
			if check(s, j+1, x-y) {
				return true
			}
		}
		return false
	}
	for i := 1; i <= n; i++ {
		x := i * i
		s := strconv.Itoa(x)
		if check(s, 0, i) {
			ans += x
		}
	}
	return
}

func main() {
	testCases := []struct {
		n int
	}{
		{n: 100},
	}
	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.n)
		fmt.Printf("OUTPUT: %v\n\n", punishmentNumber(testCase.n))
	}
}
