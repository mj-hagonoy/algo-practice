package main

import "fmt"

func main() {
	r := Solution([]int{10, 2, 5, 1, 8, 20})
	fmt.Printf("result %d\n", r)
	r = Solution([]int{10, 50, 5, 1})
	fmt.Printf("result %d\n", r)
}

func Solution(A []int) int {
	if len(A) < 3 {
		return 0
	}

	A = _sort(A)
	for i := 0; i < len(A)-2; i++ {
		p, q, r := i, i+1, i+2
		if A[p]+A[q] > A[r] && A[q]+A[r] > A[p] && A[r]+A[p] > A[q] {
			return 1
		}
	}
	return 0
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
Task Score: 100%
An array A consisting of N integers is given. A triplet (P, Q, R) is triangular if 0 ≤ P < Q < R < N and:

A[P] + A[Q] > A[R],
A[Q] + A[R] > A[P],
A[R] + A[P] > A[Q].
For example, consider array A such that:

  A[0] = 10    A[1] = 2    A[2] = 5
  A[3] = 1     A[4] = 8    A[5] = 20
Triplet (0, 2, 4) is triangular.

Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers, returns 1 if there exists a triangular triplet for this array and returns 0 otherwise.

For example, given array A such that:

  A[0] = 10    A[1] = 2    A[2] = 5
  A[3] = 1     A[4] = 8    A[5] = 20
the function should return 1, as explained above. Given array A such that:

  A[0] = 10    A[1] = 50    A[2] = 5
  A[3] = 1
the function should return 0.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].

*/
