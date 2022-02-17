package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) //0,1
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      //1,2
	fmt.Println(twoSum([]int{3, 3}, 6))         //0,1
}

func twoSum(nums []int, target int) []int {
	diff := make(map[int]int)
	numPos := make(map[int][]int)
	for i, val := range nums {
		diff[val] = target - val
		numPos[val] = append(numPos[val], i)
	}
	for val1, val2 := range diff {
		positions1 := numPos[val1]
		positions2, ok := numPos[val2]
		if !ok {
			continue
		}

		for _, pos1 := range positions1 {
			for _, pos2 := range positions2 {
				if pos1 != pos2 {
					return []int{pos1, pos2}
				}
			}
		}

	}
	return nil
}

/*
TWO SUMS
=======
Result: ACCEPTED
Runtime: 12 ms, faster than 49.02% of Go online submissions for Two Sum.
Memory Usage: 7.8 MB, less than 5.67% of Go online submissions for Two Sum.
====

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

*/
