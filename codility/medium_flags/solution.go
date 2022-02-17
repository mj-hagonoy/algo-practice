package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Solution([]int{5}))                                  //expected 0
	fmt.Println(Solution([]int{5, 10, 4}))                           //expected 1
	fmt.Println(Solution([]int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2})) //expected 3
	fmt.Println(Solution([]int{3, 2, 1}))                            //expected 0
	fmt.Println(Solution([]int{0, 0, 0, 0, 0, 1, 0, 1, 0, 1}))       //expected 2

}

func Solution(A []int) int {
	if len(A) < 3 {
		return 0
	}

	//find peaks, and distance from prev peak
	var peaks []int
	var dist []int
	prevPeak := 0
	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			//this is a peak
			peaks = append(peaks, i)
			dist = append(dist, i-prevPeak)
			prevPeak = i
		}
	}
	N := len(peaks)
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1 //can only set 1 flag
	}
	max := 1
	s := int(math.Sqrt(float64(len(A))))
	for s > 0 {
		max = _max(max, getMaxFlags(dist, s))
		s--
	}

	return max
}

func getMaxFlags(dist []int, f int) int {
	flag := 1 //always put flag in peak 0
	k := f
	dist_from_prev_flag := 0
	for i := 1; i < len(dist) && k > 0; i++ {
		dist_from_prev_flag += dist[i]
		if dist_from_prev_flag >= f {
			k--
			flag++
			dist_from_prev_flag = 0 //reset
		}
	}
	if flag > f {
		return 0 //not fit
	}
	return flag
}

func _max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
TASK SCORE: 86%
A non-empty array A consisting of N integers is given.

A peak is an array element which is larger than its neighbours. More precisely, it is an index P such that 0 < P < N − 1 and A[P − 1] < A[P] > A[P + 1].

For example, the following array A:

    A[0] = 1
    A[1] = 5
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2
has exactly four peaks: elements 1, 3, 5 and 10.

You are going on a trip to a range of mountains whose relative heights are represented by array A, as shown in a figure below. You have to choose how many flags you should take with you. The goal is to set the maximum number of flags on the peaks, according to certain rules.



Flags can only be set on peaks. What's more, if you take K flags, then the distance between any two flags should be greater than or equal to K. The distance between indices P and Q is the absolute value |P − Q|.

For example, given the mountain range represented by array A, above, with N = 12, if you take:

two flags, you can set them on peaks 1 and 5;
three flags, you can set them on peaks 1, 5 and 10;
four flags, you can set only three flags, on peaks 1, 5 and 10.
You can therefore set a maximum of three flags in this case.

Write a function:

func Solution(A []int) int

that, given a non-empty array A of N integers, returns the maximum number of flags that can be set on the peaks of the array.

For example, the following array A:

    A[0] = 1
    A[1] = 5
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2
the function should return 3, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..400,000];
each element of array A is an integer within the range [0..1,000,000,000].


*/
