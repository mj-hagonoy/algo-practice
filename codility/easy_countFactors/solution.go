package main

import "fmt"

func main() {
	fmt.Println(Solution(4))  //expected 3
	fmt.Println(Solution(5))  //expected 2
	fmt.Println(Solution(24)) //expected 8
}

func Solution(N int) int {
	if N == 1 {
		return 1
	}
	factors := 2 //N & 1
	i := 2
	for (i * i) < N {
		if N%i == 0 {
			factors += 2
		}
		i++
	}
	if i*i == N {
		factors++
	}
	return factors
}

/*
TASK SCORE 100%
A positive integer D is a factor of a positive integer N if there exists an integer M such that N = D * M.

For example, 6 is a factor of 24, because M = 4 satisfies the above condition (24 = 6 * 4).

Write a function:

func Solution(N int) int

that, given a positive integer N, returns the number of its factors.

For example, given N = 24, the function should return 8, because 24 has 8 factors, namely 1, 2, 3, 4, 6, 8, 12, 24. There are no other factors of 24.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..2,147,483,647].

*/
