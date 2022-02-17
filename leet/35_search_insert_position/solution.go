package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) //expected 2
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) //expected 1
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7)) //expected 4
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0)) //expected 0
}

func searchInsert(nums []int, target int) int {
	pos := 0
	for i, val := range nums {
		if val >= target {
			return i
		}
		pos++
	}

	return pos
}

/*
SEARCH INSERT POSITION
====
Result: ACCEPTED
Runtime: 4 ms, faster than 80.87% of Go online submissions for Search Insert Position.
Memory Usage: 3 MB, less than 50.31% of Go online submissions for Search Insert Position.
====
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4


Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104
*/
