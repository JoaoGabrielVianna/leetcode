# 1726. Tuple with Same Product

**`Medium`**

Given an array nums of distinct positive integers, return the number of tuples (a, b, c, d) such that a * b = c * d where a, b, c, and d are elements of nums, and a != b != c != d.

### Example 1:

**Input:** nums = [2,3,4,6]
**Output:** 8

**Explanation:** 

There are 8 valid tuples:
(2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3)
(3,4,2,6) , (4,3,2,6) , (3,4,6,2) , (4,3,6,2)
Example 2:

**Input:** nums = [1,2,4,5,10]

**Output:** 16

**Explanation:** 

There are 16 valid tuples:
(1,10,2,5) , (1,10,5,2) , (10,1,2,5) , (10,1,5,2)
(2,5,1,10) , (2,5,10,1) , (5,2,1,10) , (5,2,10,1)
(2,10,4,5) , (2,10,5,4) , (10,2,4,5) , (10,2,5,4)
(4,5,2,10) , (4,5,10,2) , (5,4,2,10) , (5,4,10,2)
 

### Constraints:

- `1 <= nums.length <= 1000`
- `1 <= nums[i] <= 104`
- `All elements in nums are distinct.`

### Solution

```go
func tupleSameProduct(nums []int) int {
	cnt := map[int]int{}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			x := nums[i] * nums[j]
			cnt[x]++
		}
	}
	ans := 0
	for _, v := range cnt {
		ans += v * (v - 1) / 2
	}
	return ans << 3
}


```

### Test

```go
func main() {

	testCases := []struct {
		nums []int
	}{
		{nums: []int{2, 3, 4, 6}},
		{nums: []int{1, 2, 4, 5, 10}},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.nums)
		fmt.Printf("OUTPUT: %v\n\n", tupleSameProduct(testCase.nums))
	}
}

```