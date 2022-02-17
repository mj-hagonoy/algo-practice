//https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/
//solution OK 100%
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution([]int{2, 3, 1, 5}))
	fmt.Println(Solution([]int{2}))
}

func Solution(A []int) int {
	if len(A) == 0 {
		return 1
	}

	n := len(A) + 1
	r := n * (n + 1) / 2
	sum := 0
	for _, val := range A {
		sum += val
	}

	return r - sum
}
