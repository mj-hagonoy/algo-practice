package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	v := 0
	x1 := x
	for x1 > 0 {
		v = (v * 10) + (x1 % 10)
		x1 = x1 / 10
	}
	return x == v
}

/*
PALINDROME NUMBER
===
Result: ACCEPTED
Runtime: 43 ms, faster than 12.90% of Go online submissions for Palindrome Number.
Memory Usage: 5.2 MB, less than 72.93% of Go online submissions for Palindrome Number.
=====

Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.

For example, 121 is a palindrome while 123 is not.


Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


Constraints:

-231 <= x <= 231 - 1


Follow up: Could you solve it without converting the integer to a string?
*/
