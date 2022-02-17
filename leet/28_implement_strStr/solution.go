package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))      //expected 2
	fmt.Println(strStr("aaaaaaaaaa", "bb")) //expected -1
	fmt.Println(strStr("", ""))             //expected 0
	fmt.Println(strStr("abc", "c"))         //expected 2
}

func strStr(haystack string, needle string) int {
	if needle == haystack || needle == "" {
		return 0
	}

	nn := len(needle) - 1
	nh := len(haystack)
	for i := 0; i < nh; i++ {
		if i+nn < nh && haystack[i:i+nn+1] == needle {
			return i
		}
	}

	return -1
}

/*
IMPLEMENT strStr()
====
Result: ACCEPTED
Runtime: 4 ms, faster than 84.73% of Go online submissions for Implement strStr().
Memory Usage: 2.5 MB, less than 72.85% of Go online submissions for Implement strStr().
=====
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().



Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Example 3:

Input: haystack = "", needle = ""
Output: 0


Constraints:

0 <= haystack.length, needle.length <= 5 * 104
haystack and needle consist of only lower-case English characters.

*/
