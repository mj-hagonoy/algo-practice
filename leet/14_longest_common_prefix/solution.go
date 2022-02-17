package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) //expected "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    //expected ""
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = getCommonPrefix(prefix, strs[i])
		if prefix == "" {
			return prefix
		}
	}
	return prefix
}

func getCommonPrefix(str1 string, str2 string) string {
	if len(str2) < len(str1) {
		str1, str2 = str2, str1
	}

	prefix := ""
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return prefix
		}

		prefix = prefix + string(str1[i])
	}
	return prefix
}

/*
LONGEST COMMON PREFIX
====
Result: ACCEPTED
Runtime: 4 ms, faster than 25.49% of Go online submissions for Longest Common Prefix.
Memory Usage: 5.4 MB, less than 7.71% of Go online submissions for Longest Common Prefix.
====
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lower-case English letters.

*/
