package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("ab")) //expected 1
}

func lengthOfLastWord(s string) int {
	end := len(s) - 1
	for end >= 0 && rune(s[end]) == ' ' {
		end--
	}
	start := end

	for start >= 0 && rune(s[start]) != ' ' {
		start--
	}

	return end - start
}

/*
58. LENGTH OF LAST WORD
====
Result: ACCEPTED
Runtime: 0 ms, faster than 100.00% of Go online submissions for Length of Last Word.
Memory Usage: 2.1 MB, less than 87.42% of Go online submissions for Length of Last Word.
====
Given a string s consisting of some words separated by some number of spaces, return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.



Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.


Constraints:

1 <= s.length <= 104
s consists of only English letters and spaces ' '.
There will be at least one word in s.

*/
