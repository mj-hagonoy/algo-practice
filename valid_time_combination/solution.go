package main

import (
	"fmt"
)

/*
Task: Given a list of 4 integers, find all possible valid 24 hour times (eg: 12:34) that the given integers can be assembled into and return the total number of valid times.
You can not use the same number twice.
Times such as 34:12 and 12:60 are not valid.

Notes: Input integers can not be negative.
Input integers can yeald 0 possible valid combinations.

Example:
	Input: [1, 2, 3, 4]
	Valid times: ["12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "23:14", "23:41"]
	Return: 8
*/
func possibleTimes(digits []int) int {
	// Your code here.
	combinations := make(map[int]bool)
	getCombinations(digits, 4, combinations)
	return len(combinations)
}

func main() {
	// Example test cases.
	testCases := [][]int{
		{1, 2, 3, 4}, // 10, including 21:43, 21:34
		{2, 2, 1, 9}, // 4
		{2, 2, 5, 9}, // 1
		{0, 0, 0, 0}, // 1
		{2, 4, 0, 0}, // 8 20:04, 20:40, 00:24, 00:42, 02:04, 02:40, 04:02, 04:20
		{9, 9, 9, 9}, // 0
	}

	for _, input := range testCases {
		fmt.Println(possibleTimes(input))
	}
}

// isValidHour checks if hhmm is within (00:00,24:00)
func isValidHour(hhmm int) bool {
	if !(hhmm >= 0 && hhmm < 2400) {
		return false
	}
	if (hhmm % 100) >= 60 {
		return false
	}
	return true
}

// getNumCombinations generates valid time combinations
func getCombinations(digits []int, size int, found map[int]bool) {
	if size == 1 {
		hhmm := (digits[0] * 1000) + (digits[1] * 100) + (digits[2] * 10) + digits[3]
		if _, ok := found[hhmm]; !ok {
			if isValidHour(hhmm) {
				found[hhmm] = true
			}
		}

		return
	}

	for i := 0; i < size; i++ {
		getCombinations(digits, size-1, found)
		if size%2 == 1 {
			digits[0], digits[size-1] = digits[size-1], digits[0]
		} else {
			digits[i], digits[size-1] = digits[size-1], digits[i]
		}
	}
}

//Go Playground: https://go.dev/play/p/N_CYoRFoDO3
