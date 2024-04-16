# 2667. Create Hello World Function
### Write a function createHelloWorld. It should return a new function that always returns "Hello World". 
* **`Easy`**

### Exemple 1:
> **Input:** args = []<br/>
> **Output:** "Hello World"<br/>
> **Explanation:**<br/>
> const f = createHelloWorld();<br/>
> f( ); // "Hello World"
>
> The function returned by createHelloWorld should always return "Hello World".

### Exemple 1:
> **Input:** args = [{ },null,42]<br/>
> **Output:** "Hello World"<br/>
> **Explanation:**<br/>
> const f = createHelloWorld();<br/>
> f({ }, null, 42); // "Hello World"
>
> Any arguments could be passed to the function but it should still always return "Hello World".


**Constraints**:

`0 <= args.length <= 10`

#
# My answer:
```js
/*** This Function allways return "Hello World"
 * @returns `"Hello World"`
 */
var createHelloWorld = function () {
    return function () {
        return "Hello World"
    }
}

const answer = function () {
    console.log(createHelloWorld()());
}

answer(); // "Hello World"
```

