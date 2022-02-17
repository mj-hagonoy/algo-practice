package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{4, 9}))       //50
	fmt.Println(plusOne([]int{9, 9}))       //100
	fmt.Println(plusOne([]int{4, 3, 2, 1})) //4322
}

func plusOne(digits []int) []int {
	n := len(digits)
	if digits[n-1] < 9 {
		digits[n-1] = digits[n-1] + 1
		return digits
	} else {
		result := make([]int, 1)
		result = append(result, digits...)

		carryOver := 1 //9 + 1 = 10
		result[n] = 0
		n--
		for carryOver > 0 {
			sum := result[n] + carryOver
			carryOver = sum / 10
			result[n] = sum % 10
			n--
		}
		if result[0] == 0 {
			return result[1:]
		}
		return result
	}
}

/*
66 PLUS ONE
====
Result: ACCEPTED
Runtime: 0 ms, faster than 100.00% of Go online submissions for Plus One.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Plus One.
====
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.



Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].


Constraints:

1 <= digits.length <= 100
0 <= digits[i] <= 9
digits does not contain any leading 0's.
*/
