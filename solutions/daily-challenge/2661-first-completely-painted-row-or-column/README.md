# 2661. First Completely Painted Row or Column

**`Medium`**

You are given a 0-indexed integer array `arr`, and an `m x n` integer matrix `mat`. `arr` and `mat` both contain all the integers in the range `[1, m * n]`.

Go through each index `i` in `arr` starting from index `0` and paint the cell in `mat` containing the integer `arr[i]`.

Return the smallest index i at which either a row or a column will be completely painted in `mat`.

### Example 1:

image explanation for example 1
**Input:** arr = [1,3,4,2], mat = [[1,4],[2,3]]

**Output:** 2

**Explanation:**

The moves are shown in order, and both the first row and second column of the matrix become fully painted at arr[2].
### Example 2:

image explanation for example 2
**Input:** arr = [2,8,7,4,1,3,5,6,9], mat = [[3,2,5],[1,4,6],[8,7,9]]

**Output:** 3

**Explanation:** 

The second column becomes fully painted at arr[3].
 

### Constraints:

- `m == mat.length`
- `n = mat[i].length`
- `arr.length == m * n`
- `1 <= m, n <= 105`
- `1 <= m * n <= 105`
- `1 <= arr[i], mat[r][c] <= m * n`
- All the integers of `arr` are unique.
- All the integers of `mat` are unique.

### Solution

```go
func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	idx := map[int][2]int{}
	for i := range mat {
		for j := range mat[i] {
			idx[mat[i][j]] = [2]int{i, j}
		}
	}
	row := make([]int, m)
	col := make([]int, n)
	for k := 0; ; k++ {
		x := idx[arr[k]]
		i, j := x[0], x[1]
		row[i]++
		col[j]++
		if row[i] == n || col[j] == m {
			return k
		}
	}
}
```

### Test 
```go
func main() {
	testCases := []struct {
		arr []int
		mat [][]int
	}{
		{arr: []int{1, 3, 4, 2}, mat: [][]int{{1, 4}, {2, 3}}},
		{arr: []int{2, 8, 7, 4, 1, 3, 5, 6, 9}, mat: [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT ARR:", testCase.arr)
		fmt.Println("INPUT MAT:", testCase.mat)
		fmt.Printf("OUTPUT: %d\n\n", firstCompleteIndex(testCase.arr, testCase.mat))

	}
}

```