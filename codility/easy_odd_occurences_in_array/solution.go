//https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/
//Solution NG 100% (N: 999_999 exceeds 6s)
package main

import (
	"codility/inputreader"
	"fmt"
)

func main() {
	input := inputreader.GetInputFromFileAsInt("input.txt")
	for i, line := range input {
		fmt.Printf("[%d]\t input: %+v \t result: %+v\n", i, line, Solution(line))
	}
}

func Solution(A []int) int {
	pairs := make(map[int]int, (len(A)/2)+1)
	for _, val := range A {
		pairs[val]++
	}

	for i, val := range pairs {
		if val%2 != 0 {
			return i
		}
	}
	return -1
}
