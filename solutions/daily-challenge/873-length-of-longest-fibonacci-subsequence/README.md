# 873. Length of Longest Fibonacci Subsequence
**`Medium`**

A sequence x1, x2, ..., xn is Fibonacci-like if:

n >= 3
xi + xi+1 == xi+2 for all i + 2 <= n
Given a strictly increasing array arr of positive integers forming a sequence, return the length of the longest Fibonacci-like subsequence of arr. If one does not exist, return 0.

A subsequence is derived from another sequence arr by deleting any number of elements (including none) from arr, without changing the order of the remaining elements. For example, [3, 5, 8] is a subsequence of [3, 4, 5, 6, 7, 8].

 

### Example 1:

**Input:** arr = [1,2,3,4,5,6,7,8]

**Output:** 5

**Explanation:** 

The longest subsequence that is fibonacci-like: [1,2,3,5,8].

### Example 2:

**Input:** arr = [1,3,7,11,12,14,18]

**Output:** 3

**Explanation:** 

The longest subsequence that is fibonacci-like: [1,11,12], [3,11,14] or [7,11,18].
 

### Constraints:

- `3 <= arr.length <= 1000`
- `1 <= arr[i] < arr[i + 1] <= 109`

### Solution

```go
func lenLongestFibSubseq(arr []int) (ans int) {
	n := len(arr)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}

	d := make(map[int]int)
	for i, x := range arr {
		d[x] = i
		for j := 0; j < i; j++ {
			f[i][j] = 2
		}
	}

	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			t := arr[i] - arr[j]
			if k, ok := d[t]; ok && k < j {
				f[i][j] = max(f[i][j], f[j][k]+1)
				ans = max(ans, f[i][j])
			}
		}
	}

	return
}
```

### Test

```go
func main() {
	testCases := []struct {
		arr []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{1, 3, 7, 11, 13, 14, 16}},
	}
	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.arr)
		fmt.Println("OUTPUT:", lenLongestFibSubseq(testCase.arr))
	}
}

```