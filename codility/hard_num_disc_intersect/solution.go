package main

import (
	"fmt"
)

func main() {
	r := Solution([]int{1, 5, 2, 1, 4, 0})
	if r == 11 {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("FAIL: got %d", r)
	}
}

func Solution(A []int) int {
	opening := make([]int, len(A))
	closing := make([]int, len(A))
	for pos, val := range A {
		open, close := pos-val, pos+val
		opening[pos] = open
		closing[pos] = close
	}

	opening = _sort(opening)
	closing = _sort(closing)

	cnt := 0
	j := 0
	for i := 0; i < len(A); i++ {
		for j < len(A) && closing[i] >= opening[j] {
			cnt += j - i
			j++
		}
		if cnt > 10000000 {
			return -1
		}
	}

	return cnt
}

func _sort(input []int) []int {
	n := len(input)
	if n < 2 {
		return input
	}
	if n == 2 {
		if input[0] > input[1] {
			input[1], input[0] = input[0], input[1]
		}
		return input
	}

	mid := n / 2
	left := _sort(input[0:mid])
	right := _sort(input[mid:])

	return _mergeSort(left, right)
}

func _mergeSort(a []int, b []int) []int {
	n_a := len(a)
	n_b := len(b)
	i := 0
	j := 0
	x := 0
	var merged []int
	for i < n_a && j < n_b {
		if a[i] <= b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
		x++
	}

	if i < n_a {
		merged = append(merged, a[i:]...)
	}
	if j < n_b {
		merged = append(merged, b[j:]...)
	}
	return merged
}

/*
TASK SCORE : 100%
We draw N discs on a plane. The discs are numbered from 0 to N − 1. An array A of N non-negative integers, specifying the radiuses of the discs, is given. The J-th disc is drawn with its center at (J, 0) and radius A[J].

We say that the J-th disc and K-th disc intersect if J ≠ K and the J-th and K-th discs have at least one common point (assuming that the discs contain their borders).

The figure below shows discs drawn for N = 6 and A as follows:

  A[0] = 1
  A[1] = 5
  A[2] = 2
  A[3] = 1
  A[4] = 4
  A[5] = 0


There are eleven (unordered) pairs of discs that intersect, namely:

discs 1 and 4 intersect, and both intersect with all the other discs;
disc 2 also intersects with discs 0 and 3.
Write a function:

func Solution(A []int) int

that, given an array A describing N discs as explained above, returns the number of (unordered) pairs of intersecting discs. The function should return −1 if the number of intersecting pairs exceeds 10,000,000.

Given array A shown above, the function should return 11, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [0..2,147,483,647].


*/
