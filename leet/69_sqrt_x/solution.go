package main

import "fmt"

func main() {
	fmt.Println(mySqrt(856035737)) // 29258
	fmt.Println(mySqrt(1))         //1
	fmt.Println(mySqrt(9))         //3
	fmt.Println(mySqrt(2))         // 1
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	//binary
	start, end := 1, x

	for start <= end {
		root := start + (end-start)/2
		if root*root > x {
			end = root - 1
		} else if (root+1)*(root+1) > x {
			return root
		} else {
			start = root + 1
		}
	}
	return start
}

/*
69 SQRT(X)
===
Result: ACCEPTED
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sqrt(x).
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Sqrt(x).
===
Given a non-negative integer x, compute and return the square root of x.

Since the return type is an integer, the decimal digits are truncated, and only the integer part of the result is returned.

Note: You are not allowed to use any built-in exponent function or operator, such as pow(x, 0.5) or x ** 0.5.



Example 1:

Input: x = 4
Output: 2
Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since the decimal part is truncated, 2 is returned.


Constraints:

0 <= x <= 231 - 1
*/
