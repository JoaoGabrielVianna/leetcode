package main

import "fmt"

type ProductOfNumbers struct {
	s []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{[]int{1}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.s = []int{1}
		return
	}
	this.s = append(this.s, this.s[len(this.s)-1]*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.s)
	if n <= k {
		return 0
	}
	return this.s[len(this.s)-1] / this.s[len(this.s)-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */

func main() {
	testCases := []struct {
		operations []string
	}{
		{operations: []string{"ProductOfNumbers", "add", "add", "add", "add", "add", "getProduct", "add", "getProduct"}},
	}

	for i, tc := range testCases {
		fmt.Println("INPUT:", i+1, tc.operations)
		fmt.Printf("OUTPUT: %v\n\n", Constructor())
	}
}
