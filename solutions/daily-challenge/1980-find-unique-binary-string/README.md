# 1980. Find Unique Binary String
Solved
**`Medium`**

Given an array of strings nums containing n unique binary strings each of length n, return a binary string of length n that does not appear in nums. If there are multiple answers, you may return any of them.

### Example 1:

**Input:** nums = ["01","10"]

**Output:** "11"

**Explanation:** 

"11" does not appear in nums. "00" would also be correct.

### Example 2:

**Input:** nums = ["00","01"]

**Output:** "11"

**Explanation:** 

"11" does not appear in nums. "10" would also be correct.

### Example 3:

**Input:** nums = ["111","011","001"]

**Output:** "101"

**Explanation:** 

"101" does not appear in nums. "000", "010", "100", and "110" would also be correct.
 

### Constraints:

- `n == nums.length`
- `1 <= n <= 16`
- `nums[i].length == n`
- `nums[i] is either '0' or '1'.`
- `All the strings of nums are unique.`

### Solution

```go
func findDifferentBinaryString(nums []string) string {
	mask := 0
	for _, x := range nums {
		mask |= 1 << strings.Count(x, "1")
	}
	for i := 0; ; i++ {
		if mask>>i&1 == 0 {
			return strings.Repeat("1", i) + strings.Repeat("0", len(nums)-i)
		}
	}
}

```

### Test

```go
func main() {
	testCases := []struct {
		nums []string
	}{
		{nums: []string{"01", "10", "11", "00"}},
		{nums: []string{"00", "01", "11", "10"}},
		{nums: []string{"111", "011", "001", "101", "100", "010", "110", "000"}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.nums)
		fmt.Printf("OUTPUT: %v\n\n", findDifferentBinaryString(testCase.nums))
	}
}

```