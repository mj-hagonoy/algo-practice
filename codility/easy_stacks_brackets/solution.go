package main

import "fmt"

func main() {
	answer := Solution("[{()()}]")
	fmt.Println(answer)
}

func Solution(S string) int {
	stack := NewStack(len(S))
	for i := 0; i < len(S); i++ {
		val := S[i]
		if val == '{' || val == '[' || val == '(' {
			stack.push(rune(val))
			continue
		}
		//means val is })]
		if stack.isEmpty() {
			return 0
		}

		top := stack.peek()
		if val == ')' && top == '(' {
			stack.pop()
			continue
		}
		if val == '}' && top == '{' {
			stack.pop()
			continue
		}
		if val == ']' && top == '[' {
			stack.pop()
			continue
		}
		return 0
	}

	if !stack.isEmpty() {
		return 0
	}
	return 1
}

type Stack struct {
	n     int
	stack []rune
}

func (s *Stack) isEmpty() bool {
	return s.n == 0
}

func (s *Stack) push(v rune) {
	s.stack[s.n] = v
	s.n++
}

func (s *Stack) pop() rune {
	v := s.stack[s.n]
	s.n--
	return v
}

func (s *Stack) peek() rune {
	return s.stack[s.n-1]
}

func NewStack(max int) *Stack {
	return &Stack{
		n:     0,
		stack: make([]rune, max),
	}
}

/*
TASK SCORE : 100%
A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:

S is empty;
S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.
For example, the string "{[()()]}" is properly nested but "([)()]" is not.

Write a function:

func Solution(S string) int

that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.

For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..200,000];
string S consists only of the following characters: "(", "{", "[", "]", "}" and/or ")".

*/
