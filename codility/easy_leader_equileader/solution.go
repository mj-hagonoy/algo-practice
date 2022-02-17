package main

import "fmt"

func main() {
	answer := Solution([]int{4, 3, 4, 4, 4, 2})
	//expected 2
	fmt.Println(answer)
}

func Solution(A []int) int {
	cntrR := make(map[int]int)
	cntrL := make(map[int]int)

	for _, val := range A {
		if cnt, ok := cntrR[val]; ok {
			cntrR[val] = cnt + 1
		} else {
			cntrR[val] = 1
		}
	}

	cntEquiLeaders := 0
	leader := A[0]

	for i, val := range A {
		cntrR[val] = cntrR[val] - 1
		cntrL[val] = cntrL[val] + 1

		if i == 0 || cntrL[val] > cntrL[leader] {
			leader = val
		}

		if (i+1)/2 < cntrL[leader] && cntrR[leader] > (len(A)-i-1)/2 {
			cntEquiLeaders++
		}
	}

	return cntEquiLeaders
}

/*
TASK SCORE : 100%

A non-empty array A consisting of N integers is given.

The leader of this array is the value that occurs in more than half of the elements of A.

An equi leader is an index S such that 0 ≤ S < N − 1 and two sequences A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N − 1] have leaders of the same value.

For example, given array A such that:

    A[0] = 4
    A[1] = 3
    A[2] = 4
    A[3] = 4
    A[4] = 4
    A[5] = 2
we can find two equi leaders:

0, because sequences: (4) and (3, 4, 4, 4, 2) have the same leader, whose value is 4.
2, because sequences: (4, 3, 4) and (4, 4, 2) have the same leader, whose value is 4.
The goal is to count the number of equi leaders.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the number of equi leaders.

For example, given:

    A[0] = 4
    A[1] = 3
    A[2] = 4
    A[3] = 4
    A[4] = 4
    A[5] = 2
the function should return 2, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].

*/
