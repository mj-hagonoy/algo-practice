package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))      // 100
	fmt.Println(addBinary("1010", "1011")) //10101
	fmt.Println(addBinary("100", "100"))   //1000
}

func addBinary(a string, b string) string {
	max := _max(len(a), len(b))
	a = _append(a, max)
	b = _append(b, max)

	result := ""
	carryOver := 0

	for i := max - 1; i >= 0; i-- {
		sum, co := _add((int(a[i] - 48)), int(b[i]-48), carryOver)
		carryOver = co
		result = fmt.Sprintf("%d%s", sum, result)
	}

	if carryOver > 0 {
		result = fmt.Sprintf("%d%s", carryOver, result)
	}
	return result
}

func _max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func _append(s string, size int) string {
	if len(s) == size {
		return s
	}

	n := size - len(s) - 1
	for n >= 0 {
		s = "0" + s
		n--
	}
	return s
}

func _add(a int, b int, c int) (int, int) {
	return (a + b + c) % 2, (a + b + c) / 2
}

/*
67 ADD BINARY
====
Result: Accepted
Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Binary.
Memory Usage: 2.8 MB, less than 41.73% of Go online submissions for Add Binary.
====
Given two binary strings a and b, return their sum as a binary string.



Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"


Constraints:

1 <= a.length, b.length <= 104
a and b consist only of '0' or '1' characters.
Each string does not contain leading zeros except for the zero itself.
*/
