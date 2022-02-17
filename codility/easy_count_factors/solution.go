//count factors
//solution: 64% [Performance: NG]
package main

import (
	"codility/inputreader"
	"fmt"
)

func main() {
	input := inputreader.GetInputFromFileAsInt("input.txt")
	for i, line := range input {
		fmt.Printf("[%d]\t input: %+v \t result: %+v\n", i, line, Solution(line[0]))
	}
}

func Solution(N int) int {
	factors := make(map[int]bool)
	factors[N] = true
	factors[1] = true

	for f := N - 1; f >= 2; f-- {
		if factors[f] {
			continue
		}
		if N%f != 0 {
			continue
		}

		factors[f] = true
	}

	return len(factors)
}
