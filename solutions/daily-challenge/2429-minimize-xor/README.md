# 2429. Minimize XOR

**`Medium`**

Given two positive integers `num1` and `num2`, find the positive integer `x` such that:

x has the same number of set bits as `num2`, and
The value `x XOR num1` is minimal.
Note that XOR is the bitwise `XOR` operation.

Return the integer `x`. The test cases are generated such that `x` is uniquely determined.

The number of set bits of an integer is the number of `1`'s in its binary representation.

 

### Example 1:

**Input:** num1 = 3, num2 = 5
**Output:** 3
**Explanation:**

The binary representations of num1 and num2 are 0011 and 0101, respectively.
The integer 3 has the same number of set bits as num2, and the value 3 XOR 3 = 0 is minimal.
### Example 2:

**Input:** num1 = 1, num2 = 12
**Output:** 3
**Explanation:**

The binary representations of num1 and num2 are 0001 and 1100, respectively.
The integer 3 has the same number of set bits as num2, and the value 3 XOR 1 = 2 is minimal.
 

### Constraints:

- `1 <= num1, num2 <= 109`

### Soluction

```go
func minimizeXor(num1 int, num2 int) int {
	bits := 0
	for n := num2; n != 0; n >>= 1 {
		bits += n & 1
	}

	// Create bit arrays for num1 and res
	bitArray := make([]int, 32) // 32 bits to handle up to 32-bit integers
	for i := 0; i < 32; i++ {
		bitArray[i] = (num1 >> i) & 1
	}

	resArray := make([]int, 32)

	// Place `1`s in positions matching `num1` where possible
	for i := 31; i >= 0; i-- {
		if bits == 0 {
			break
		}
		if bitArray[i] == 1 {
			resArray[i] = 1
			bits--
		}
	}

	// Fill remaining bits with `1`s in lower positions
	for i := 0; i < 32; i++ {
		if bits == 0 {
			break
		}
		if resArray[i] == 0 {
			resArray[i] = 1
			bits--
		}
	}

	res := 0
	for i := 31; i >= 0; i-- {
		res = (res << 1) | resArray[i]
	}
	return res
}
```


### Test

```go
func main() {

	testCases := []struct {
		num1 int
		num2 int
	}{
		{num1: 3, num2: 5},
		{num1: 1, num2: 12},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT NUM1", testCase.num1)
		fmt.Println("INPUT NUM2", testCase.num2)
		fmt.Println("OUTPUT", minimizeXor(testCase.num1, testCase.num2))
	}
}
```