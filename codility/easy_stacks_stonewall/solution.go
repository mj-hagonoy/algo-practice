package main

import "fmt"

func main() {
	answer := Solution([]int{8, 8, 5, 7, 9, 8, 7, 4, 8})
	fmt.Println(answer)
}

func Solution(H []int) int {
	if len(H) == 0 {
		return 0
	}
	if len(H) == 1 {
		return 1
	}
	stack := NewStack(len(H) * 5)
	blocks := 0
	stack.push(H[0])

	for i := 1; i < len(H); i++ {
		if stack.peek() == H[i] {
			continue
		}

		if stack.peek() < H[i] {
			stack.push(H[i])
			continue
		}

		if stack.peek() > H[i] {
			for !stack.isEmpty() && stack.peek() > H[i] {
				blocks++
				stack.pop()
			}
		}
		if !stack.isEmpty() && stack.peek() == H[i] {
			stack.pop()
		}
		stack.push(H[i])
	}

	return blocks + stack.n
}

type Stack struct {
	n     int
	stack []int
}

func NewStack(size int) *Stack {
	return &Stack{
		n:     0,
		stack: make([]int, size),
	}
}

func (s *Stack) push(v int) {
	s.stack[s.n] = v
	s.n++
}

func (s *Stack) pop() int {
	v := s.stack[s.n-1]
	s.stack[s.n-1] = -1
	s.n--
	return v
}

func (s *Stack) peek() int {
	return s.stack[s.n-1]
}

func (s *Stack) isEmpty() bool {
	return s.n == 0
}

/*
TASK SCORE : 100%
You are going to build a stone wall. The wall should be straight and N meters long, and its thickness should be constant; however, it should have different heights in different places. The height of the wall is specified by an array H of N positive integers. H[I] is the height of the wall from I to I+1 meters to the right of its left end. In particular, H[0] is the height of the wall's left end and H[N−1] is the height of the wall's right end.

The wall should be built of cuboid stone blocks (that is, all sides of such blocks are rectangular). Your task is to compute the minimum number of blocks needed to build the wall.

Write a function:

class Solution { public int solution(int[] H); }

that, given an array H of N positive integers specifying the height of the wall, returns the minimum number of blocks needed to build it.

For example, given array H containing N = 9 integers:

  H[0] = 8    H[1] = 8    H[2] = 5
  H[3] = 7    H[4] = 9    H[5] = 8
  H[6] = 7    H[7] = 4    H[8] = 8
the function should return 7. The figure shows one possible arrangement of seven blocks.



Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array H is an integer within the range [1..1,000,000,000].

*/
