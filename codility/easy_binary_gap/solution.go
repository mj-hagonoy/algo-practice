//https://app.codility.com/programmers/lessons/1-iterations/binary_gap/
//Solution : OK 100%

package main

import (
	"codility/inputreader"
	"fmt"
	"strconv"
)

func main() {
	input := inputreader.GetInputFromFileAsInt("input.txt")
	for i, line := range input {
		fmt.Printf("[%d]\t input: %+v [%+v]\t result: %+v\n", i, line, strconv.FormatInt(int64(line[0]), 2), Solution(line[0]))
	}

	/*
		for i := 0; i <= 1000000; i++ {
			fmt.Printf("[%d]\t input: %+v [%+v]\t result: %+v\n", i, i, strconv.FormatInt(int64(i), 2), Solution(i))
		}
	*/
}

func Solution(N int) int {
	bStr := strconv.FormatInt(int64(N), 2)

	ctr := 0
	numZeros := 0
	maxGap := 0
	for i := 0; i < len(bStr); i++ {
		c := bStr[i]

		if c == '1' {
			ctr++
			if ctr >= 2 {
				ctr--
				if numZeros > maxGap {
					maxGap = numZeros
				}
			}
			numZeros = 0
		}
		if c == '0' && ctr > 0 {
			numZeros++
		}
	}

	return maxGap
}
