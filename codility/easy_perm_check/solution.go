package main

import "fmt"

func main() {
	r := Solution([]int{7, 6, 5, 4, 10, 9, 8, 3, 2, 1})
	fmt.Println(r)
}

func Solution(A []int) int {
	n := len(A)
	r := n * (n + 1) / 2
	sum := 0
	exists := make(map[int]bool)
	for _, val := range A {
		if exists[val] {
			return 0
		}
		exists[val] = true
		sum += val
		if sum > r {
			return 0
		}
	}
	if sum == r {
		return 1
	}
	return 0
}

/**
RESULT : 100%
A non-empty array A consisting of N integers is given.

A permutation is a sequence containing each element from 1 to N once, and only once.

For example, array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
is a permutation, but array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
is not a permutation, because value 2 is missing.

The goal is to check whether array A is a permutation.

Write a function:

func Solution(A []int) int

that, given an array A, returns 1 if array A is a permutation and 0 if it is not.

For example, given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
the function should return 1.

Given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
the function should return 0.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [1..1,000,000,000].


*/
