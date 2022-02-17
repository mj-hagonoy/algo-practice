//Tape equilibrium
//Result  100%
package main

import (
	"codility/inputreader"
	"fmt"
	"math"
)

func main() {

	input := inputreader.GetInputFromFileAsInt("input.txt")
	for i, line := range input {
		fmt.Printf("[%d]\t input: %+v \t result: %+v\n", i, line, Solution(line))
	}
}

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}

	sums := make([]float64, len(A))
	for i, val := range A {
		if i == 0 {
			sums[i] = float64(val)
		} else {
			sums[i] = sums[i-1] + float64(val)
		}
	}
	if len(sums) == 1 {
		return int(math.Abs(sums[0]))
	}

	last := len(sums) - 1
	left := sums[0]
	right := sums[last] - left
	minDiff := math.Abs(left - right)
	for i := 1; i < last; i++ {
		left = sums[i]
		right = sums[last] - left
		diff := math.Abs(left - right)
		if diff < minDiff {
			minDiff = diff
		}

	}
	return int(minDiff)
}
