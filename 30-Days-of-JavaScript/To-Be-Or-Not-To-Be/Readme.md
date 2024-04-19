# 2704. To Be Or Not To Be

### Write a function `expect` that helps developers test their code. It should take in any value `val` and return an object with the following two functions.
* **`Easy`**<br/><br/>

- `toBe(val)` accepts another value and returns `true` if the two values `===`each other. If they are not equal, it should throw an error `"Not Equal"`.
- `notToBe(val)` accepts another value and returns `true` if the two values `!==` each other. If they are equal, it should throw an error `"Equal"`.

### Example 1:

> **Input:** func = () => expect(5).toBe(5)<br/>
> **Output:** {"value": true}<br/>
> **Explanation:** 5 === 5 so this expression returns true.

### Example 2:

> **Input:** func = () => expect(5).toBe(null)<br/>
> **Output:** {"error": "Not Equal"}<br/>
> **Explanation:** 5 !== null so this expression throw the error "Not Equal".

### Example 3:

> **Input:** func = () => expect(5).notToBe(null)<br/>
> **Output:** {"value": true}<br/>
> **Explanation:** 5 !== null so this expression returns true.

#
# My Solution:
```js
/**
 * This function compares the provided value with the value passed to the toBe and notToBe functions.
 * 
 * @param {*} val The value to be compared with the value passed to the toBe and notToBe functions.
 * @returns {object} An object containing the toBe and notToBe functions.
 */
var expect = function (val) {
    return {
        /**
         * Compares the provided value with the given value. 
         * 
         * @param {*} val2 The value to compare with the original value.
         * @throws {Error} Throws an error with message "Not Equal" if the values are not equal.
         * @returns {boolean} Returns true if the values are equal.
         */
        toBe: (val2) => {
            if (val !== val2) {
                throw new Error("Not Equal");

            } else {
                return true
            }

        },
        /**
        * Compares the provided value with the given value. 
        * 
        * @param {*} val3 The value to compare with the original value.
        * @throws {Error} Throws an error with message "Not Equal" if the values are equal.
        * @returns {boolean} Returns true if the values are not equal.
        */
        notToBe: (val3) => {
            if (val === val3) {
                throw new Error("Not Equal");
            } else {
                return true
            }
        }
    }
}

try {
    expect(5).toBe(null)
} catch (error) {
    console.log(error.message)
}
```