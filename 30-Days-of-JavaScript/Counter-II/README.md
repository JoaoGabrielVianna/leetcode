# 2665. Counter II

Write a function `createCounter`. It should accept an initial integer `init`. It should return an object with three functions.<br/>
- `Easy`

### The three functions are:

- `increment()` increases the current value by 1 and then returns it.
- `decrement()` reduces the current value by 1 and then returns it.
- `reset()`sets the current value to `init` and then returns it.

 

### Example 1:

> **Input:** init = 5, calls = ["increment","reset","decrement"]<br/>
> **Output:** [6,5,4]<br/>
> **Explanation:**<br/>
> const counter = createCounter(5);<br/>
> counter.increment(); // 6<br/>
> counter.reset(); // 5<br/>
> counter.decrement(); // 4

### Example 2:

> **Input:** init = 0, calls = ["increment","increment","decrement",<br/>"reset","reset"]
> **Output:** [1,2,1,0,0]<br/>
> **Explanation:**<br/>
> const counter = createCounter(0);<br/>
> counter.increment(); // 1<br/>
> counter.increment(); // 2<br/>
> counter.decrement(); // 1<br/>
> counter.reset(); // 0<br/>
> counter.reset(); // 0

 

### Constraints:

- `-1000 <= init <= 1000`
- `0 <= calls.length <= 1000`
- `calls[i]` is one of "increment", "decrement" and "reset"

#
# My Soluction:

```js
// Author: João Gabriel Vianna
// Date: 16/04/23
// source: https://leetcode.com/problems/counter-ii


/**
 * This function have 3 methods,
 * 
 * `increment() // n + 1`
 * 
 * `decrement() // n - 1`
 * 
 * `reset() // n`
 * @param {int} init 
 * @returns 
 */
var createCounter = function (init) {
    let count = init;

    function increment() {
        return ++count
    }
    function decrement() {
        return --count
    }
    function reset() {
        return (count = init)
    }

    return { increment, decrement, reset };
}


```