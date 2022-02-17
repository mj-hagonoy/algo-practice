package main

import "fmt"

func main() {
	input := []string{
		"())",
		"(()()())",
	}

	for _, s := range input {
		answer := Solution(s)
		fmt.Println(answer)
	}
}
func Solution(S string) int {
	if len(S) == 0 {
		return 1
	}

	stack := NewStack(len(S))
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack.push(rune(S[i]))
			continue
		}
		if stack.isEmpty() {
			return 0
		}
		stack.pop()
	}

	if stack.isEmpty() {
		return 1
	} else {
		return 0
	}
}

type Stack struct {
	n     int
	stack []rune
}

func NewStack(size int) *Stack {
	return &Stack{
		n:     0,
		stack: make([]rune, size),
	}
}

func (s *Stack) push(v rune) {
	s.stack[s.n] = v
	s.n++
}

func (s *Stack) pop() rune {
	v := s.stack[s.n-1]
	s.n--
	return v
}

func (s *Stack) isEmpty() bool {
	return s.n == 0
}

/*
TASK SCORE:100%

A string S consisting of N characters is called properly nested if:

S is empty;
S has the form "(U)" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.
For example, string "(()(())())" is properly nested but string "())" isn't.

Write a function:

func Solution(S string) int

that, given a string S consisting of N characters, returns 1 if string S is properly nested and 0 otherwise.

For example, given S = "(()(())())", the function should return 1 and given S = "())", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..1,000,000];
string S consists only of the characters "(" and/or ")".

*/
