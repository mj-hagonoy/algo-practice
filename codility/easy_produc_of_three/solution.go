package main

import (
	"fmt"
	"math"
)

func main() {
	c := Solution([]int{-3, 1, 2, -2, 5, 6})
	fmt.Println(c)
}

func Solution(A []int) int {
	if len(A) == 3 {
		return A[0] * A[1] * A[2]
	}
	high1 := math.MinInt64
	high2 := math.MinInt64
	high3 := math.MinInt64
	low2 := math.MinInt64
	low1 := math.MinInt64

	for _, val := range A {
		if val > high1 {
			high3 = high2
			high2 = high1
			high1 = val
		} else if val > high2 {
			high3 = high2
			high2 = val
		} else if val > high3 {
			high3 = val
		}
		if val < low1 {
			low2 = low1
			low1 = val
		} else if val < low2 {
			low2 = val
		}
	}

	p1 := high1 * high2 * high3
	p2 := high1 * low1 * low2

	if p1 > p2 {
		return p1
	}
	return p2
}

/*
Task score: 100%
A non-empty array A consisting of N integers is given. The product of triplet (P, Q, R) equates to A[P] * A[Q] * A[R] (0 ≤ P < Q < R < N).

For example, array A such that:

  A[0] = -3
  A[1] = 1
  A[2] = 2
  A[3] = -2
  A[4] = 5
  A[5] = 6
contains the following example triplets:

(0, 1, 2), product is −3 * 1 * 2 = −6
(1, 2, 4), product is 1 * 2 * 5 = 10
(2, 4, 5), product is 2 * 5 * 6 = 60
Your goal is to find the maximal product of any triplet.

Write a function:

func Solution(A []int) int

that, given a non-empty array A, returns the value of the maximal product of any triplet.

For example, given array A such that:

  A[0] = -3
  A[1] = 1
  A[2] = 2
  A[3] = -2
  A[4] = 5
  A[5] = 6
the function should return 60, as the product of triplet (2, 4, 5) is maximal.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [3..100,000];
each element of array A is an integer within the range [−1,000..1,000].

*/
