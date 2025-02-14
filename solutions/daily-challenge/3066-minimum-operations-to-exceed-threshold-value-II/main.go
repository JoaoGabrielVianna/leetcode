package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func minOperations(nums []int, k int) (ans int) {
	pq := &hp{nums}
	heap.Init(pq)
	for ; pq.Len() > 1 && pq.IntSlice[0] < k; ans++ {
		x, y := heap.Pop(pq).(int), heap.Pop(pq).(int)
		heap.Push(pq, x*2+y)
	}
	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Less(i, j int) bool { return h.IntSlice[i] < h.IntSlice[j] }
func (h *hp) Pop() interface{} {
	old := h.IntSlice
	n := len(old)
	x := old[n-1]
	h.IntSlice = old[0 : n-1]
	return x
}
func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 2, 3, 4, 5}, k: 5, want: 3},
		{nums: []int{1, 2, 3, 4, 5}, k: 6, want: 4},
		{nums: []int{1, 2, 3, 4, 5}, k: 7, want: 5},
		{nums: []int{1, 2, 3, 4, 5}, k: 8, want: 6},
	}

	for i, tc := range testCases {
		fmt.Println("INPUT:", i+1, tc.nums, tc.k)
		fmt.Printf("OUTPUT: %v\n\n", minOperations(tc.nums, tc.k))
	}
}
