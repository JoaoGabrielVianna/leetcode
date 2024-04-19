// Author: João Gabriel Vianna
// Date: 16/04/23
// source: https://leetcode.com/problems/create-hello-world-function/


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