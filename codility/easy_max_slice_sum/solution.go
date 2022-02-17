package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{1, 1, 1}))        //expected 3
	fmt.Println(Solution([]int{3, 2, -6, 4, 0})) //expected 5
}

func Solution(A []int) int {
	return maxSum(A)
}

func maxSum(input []int) int {
	sums := make([]int, len(input))

	sums[0] = input[0]
	max := input[0]

	for i := 1; i < len(input); i++ {
		sum1 := input[i]
		sum2 := input[i] + sums[i-1]
		m := _max(sum1, sum2)
		sums[i] = m
		if sums[i] > max {
			max = sums[i]
		}
	}

	return max
}

func _max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
TASK SCORE: 100%
A non-empty array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P ≤ Q < N, is called a slice of array A. The sum of a slice (P, Q) is the total of A[P] + A[P+1] + ... + A[Q].

Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers, returns the maximum sum of any slice of A.

For example, given array A such that:

A[0] = 3  A[1] = 2  A[2] = -6
A[3] = 4  A[4] = 0
the function should return 5 because:

(3, 4) is a slice of A that has sum 4,
(2, 2) is a slice of A that has sum −6,
(0, 1) is a slice of A that has sum 5,
no other slice of A has sum greater than (0, 1).
Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..1,000,000];
each element of array A is an integer within the range [−1,000,000..1,000,000];
the result will be an integer within the range [−2,147,483,648..2,147,483,647].

*/
