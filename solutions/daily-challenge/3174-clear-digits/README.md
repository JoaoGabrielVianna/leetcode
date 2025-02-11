# 3174. Clear Digits

**`Easy`**

You are given a string s.

Your task is to remove all digits by doing this operation repeatedly:

Delete the first digit and the closest non-digit character to its left.
Return the resulting string after removing all digits.

### Example 1:

**Input:** s = "abc"

**Output:** "abc"

**Explanation:**

There is no digit in the string.

### Example 2:

**Input:** s = "cb34"

**Output:** ""

**Explanation:**

First, we apply the operation on s[2], and s becomes "c4".

Then we apply the operation on s[1], and s becomes "".

 

### Constraints:

1 <= s.length <= 100
s consists only of lowercase English letters and digits.
The input is generated such that it is possible to delete all digits.

### Solution

```go
func clearDigits(s string) string {
	stk := []byte{}
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, s[i])
		}
	}
	return string(stk)
}
```

### Test

```go
func main() {
	testCases := []struct {
		s string
	}{
		{s: "abc"},
		{s: "cb34"},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.s)
		fmt.Printf("OUTPUT: %v\n\n", clearDigits(testCase.s))
	}
}

```