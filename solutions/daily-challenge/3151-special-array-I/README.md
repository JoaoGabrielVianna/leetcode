# 3151. Special Array I

**`Easy`**

An array is considered special if every pair of its adjacent elements contains two numbers with different parity.

You are given an array of integers nums. Return true if nums is a special array, otherwise, return false.

### Example 1:

**Input:** nums = [1]

**Output:** true

**Explanation:**

There is only one element. So the answer is true.

### Example 2:

**Input:** nums = [2,1,4]

**Output:** true

**Explanation:**

There is only two pairs: (2,1) and (1,4), and both of them contain numbers with different parity. So the answer is true.

### Example 3:

**Input:** nums = [4,3,1,6]

**Output:** false

**Explanation:**

nums[1] and nums[2] are both odd. So the answer is false.

### Constraints:

- `1 <= nums.length <= 100`
- `1 <= nums[i] <= 100`

### Solution

```go
func isArraySpecial(nums []int) bool {
	for i, x := range nums[1:] {
		if x%2 == nums[i]%2 {
			return false
		}
	}
	return true
}
```

### Test

```go
func main() {
	testCases := []struct {
		nums []int
	}{
		{nums: []int{1}},
		{nums: []int{2, 1, 4}},
		{nums: []int{4, 3, 1, 6}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.nums)
		fmt.Printf("OUTPUT: %v\n\n", isArraySpecial(testCase.nums))
	}
}
```