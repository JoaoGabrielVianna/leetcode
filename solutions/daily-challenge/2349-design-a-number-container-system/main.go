package main

import (
	"fmt"

	"github.com/emirpasic/gods/trees/redblacktree" // Biblioteca para Red-Black Tree
)

type NumberContainers struct {
	d map[int]int
	g map[int]*redblacktree.Tree
}

func Constructor() NumberContainers {
	return NumberContainers{map[int]int{}, map[int]*redblacktree.Tree{}}
}

func (this *NumberContainers) Change(index int, number int) {
	if oldNumber, ok := this.d[index]; ok {
		this.g[oldNumber].Remove(index)
	}
	this.d[index] = number
	if _, ok := this.g[number]; !ok {
		this.g[number] = redblacktree.NewWithIntComparator()
	}
	this.g[number].Put(index, nil)
}

func (this *NumberContainers) Find(number int) int {
	if ids, ok := this.g[number]; ok && ids.Size() > 0 {
		return ids.Left().Key.(int)
	}
	return -1
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

func main() {
	testCases := []struct {
		operations []string
		inputs     [][]int
		expected   []interface{}
	}{
		{
			operations: []string{"NumberContainers", "find", "change", "change", "change", "change", "find", "change", "find"},
			inputs:     [][]int{{}, {10}, {2, 10}, {1, 10}, {3, 10}, {5, 10}, {10}, {1, 20}, {10}},
			expected:   []interface{}{nil, -1, nil, nil, nil, nil, 1, nil, 2},
		},
	}

	for i, tc := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		var obj NumberContainers

		for j, op := range tc.operations {
			switch op {
			case "NumberContainers":
				obj = Constructor()
				fmt.Println("Constructor() -> nil")
			case "change":
				index, number := tc.inputs[j][0], tc.inputs[j][1]
				obj.Change(index, number)
				fmt.Printf("Change(%d, %d) -> nil\n", index, number)
			case "find":
				number := tc.inputs[j][0]
				result := obj.Find(number)
				fmt.Printf("Find(%d) -> %d\n", number, result)
			}
		}
		fmt.Println()
	}
}
