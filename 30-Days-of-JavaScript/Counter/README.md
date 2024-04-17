# 2620. Counter


### Given an integer `n`, return a `counter` function. This `counter` function initially returns `n` and then returns 1 more than the previous value every subsequent time it is called (`n`, `n + 1`, `n + 2,` etc).
- `Easy`
 

### Example 1:

> **Input:**<br/> 
> n = 10 <br/>
> ["call","call","call"]<br/>
> **Output:** [10,11,12]<br/>
> **Explanation:** <br/>
> counter() = 10 // The first time counter() is called, it returns n.<br/>
> counter() = 11 // Returns 1 more than the previous time.<br/>
> counter() = 12 // Returns 1 more than the previous time.

### Example 2:

> **Input:**<br/>
> n = -2<br/>
> ["call","call","call","call","call"]<br/>
> **Output:** [-2,-1,0,1,2]<br/>
> **Explanation:** counter() initially returns -2. Then increases after each sebsequent call.

**Constraints:**

`-1000 <= n <= 1000`<br/>
`0 <= calls.length <= 1000`<br/>
`calls[i] === "call"`


#
# My Solution:

```js
/**
 * This function return the current value of the counter `num`.
 * 
 *  When called again, teh returns value will be incremented by 1 `num + 1`
 * 
 * Each subsequent call to the function will returns the incremented `value`.
 * 
 * `@requires number`
 * 
 * `@returns {number}` The current value of the counter `num`
 */
var createCounter = function(n) {
    return function() {
        return n++;
    };
};


var n = 10;
var counter = createCounter(n);

console.log(counter()); // Returns: num 
console.log(counter()); // Returns: num + 1
console.log(counter()); // Returns: num + 2
```

