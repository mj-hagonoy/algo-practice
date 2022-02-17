package main

import "fmt"

func main() {
	inputs := [][]int{
		{1, 3, 6, 4, 1, 2},
		{-1, -3},
	}
	for _, input := range inputs {
		result := Solution(input)
		fmt.Printf("%+v \t %+v \n", input, result)
	}
}

func Solution(A []int) int {
	exists := make(map[int]bool)
	max := 1
	for _, val := range A {
		if val > 0 {
			exists[val] = true
			if val > max {
				max = val
			}
		}
	}

	for i := 1; i <= max; i++ {
		if !exists[i] {
			return i
		}
	}
	return max + 1
}

/* RESULT : 100%

Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].

*/
