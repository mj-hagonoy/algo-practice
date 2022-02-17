package main

import (
	"fmt"
)

type InputValue struct {
	X int
	A []int
}

func main() {
	inputs := []InputValue{
		{X: 5, A: []int{1, 3, 1, 4, 2, 3, 4, 4, 1, 5}},
		{X: 5, A: []int{1}},
		{X: 5, A: []int{5}},
		{X: 2, A: []int{2, 2}},
	}

	for i, input := range inputs {
		fmt.Printf("[%d] \t input: [%+v] \t result: [%+v]\n", i, input, Solution(input.X, input.A))
	}
}

//var max int = 100001

type Leaf struct {
	Sec    int
	Exists bool
}

func Solution(X int, A []int) int {
	if X > len(A) {
		return -1 //not enough leaves
	}
	leaves := make(map[int]Leaf, X+1) //only need leaves for path until X

	for s, l := range A {
		if l > X {
			continue
		}
		leaf := leaves[l]
		if !leaf.Exists || (leaf.Exists && s < leaf.Sec) {
			leaves[l] = Leaf{Sec: s, Exists: true}
		}
	}
	if len(leaves) != X {
		return -1
	}
	return leaves[X].Sec
}
