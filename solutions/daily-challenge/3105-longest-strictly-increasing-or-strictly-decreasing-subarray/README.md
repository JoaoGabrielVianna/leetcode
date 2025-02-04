# 3105. Longest Strictly Increasing or Strictly Decreasing Subarray

**`Easy`**

You are given an array of integers nums. Return the length of the longest 
subarray
 of nums which is either 
strictly increasing
 or 
strictly decreasing
.

### Example 1:

**Input:** nums = [1,4,3,3,2]

**Output:** 2

**Explanation:**

The strictly increasing subarrays of nums are [1], [2], [3], [3], [4], and [1,4].

The strictly decreasing subarrays of nums are [1], [2], [3], [3], [4], [3,2], and [4,3].

Hence, we return 2.

### Example 2:

**Input:** nums = [3,3,3,3]

**Output:** 1

**Explanation:**

The strictly increasing subarrays of nums are [3], [3], [3], and [3].

The strictly decreasing subarrays of nums are [3], [3], [3], and [3].

Hence, we return 1.

### Example 3:

**Input:** nums = [3,2,1]

**Output:** 3

**Explanation:**

The strictly increasing subarrays of nums are [3], [2], and [1].

The strictly decreasing subarrays of nums are [3], [2], [1], [3,2], [2,1], and [3,2,1].

Hence, we return 3.

 

### Constraints:

- `1 <= nums.length <= 50`
- `1 <= nums[i] <= 50`

### Solution

```go
func longestMonotonicSubarray(nums []int) int {
	ans := 1
	t := 1
	for i, x := range nums[1:] {
		if nums[i] < x {
			t++
			ans = max(ans, t)
		} else {
			t = 1
		}
	}
	t = 1
	for i, x := range nums[1:] {
		if nums[i] > x {
			t++
			ans = max(ans, t)
		} else {
			t = 1
		}
	}
	return ans
}
```

### Test

```go
func main() {

	testCases := []struct {
		nums []int
	}{
		{nums: []int{1, 4, 3, 3, 2}},
		{nums: []int{3, 3, 3, 3}},
		{nums: []int{3, 2, 1}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.nums)
		fmt.Printf("OUTPUT: %v\n\n", longestMonotonicSubarray(testCase.nums))
	}
}

```