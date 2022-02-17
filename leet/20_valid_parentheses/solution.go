package main

import "fmt"

func main(){
	fmt.Println(isValid("()")) //true
	fmt.Println(isValid("()[]{}")) //true
	fmt.Println(isValid("(]")) //false
}

func isValid(s string) bool {
    closers := map[rune]rune{
        '(' : ')',
        '{' : '}',
        '[' : ']',
    }
    
    
    stack := NewStack(len(s))
    for _, c := range s {
        if isOpen(c) {
            stack.push(c)
            continue
        }
        
        //c is ), } or ]
        if stack.isEmpty() {
            return false
        }
        
        if closers[stack.pop()] != c {
            return false
        }
    }
    
    return stack.isEmpty()
    
}

func isOpen(c rune) bool {
    return c == '{' || c == '[' || c == '('
}


type Stack struct{
    _stack []rune
    _size int
}

func NewStack(size int) *Stack{
    return &Stack{
        _stack : make([]rune, size),
        _size: 0,
    }
}

func (s *Stack) push(v rune) {
    s._stack[s._size] = v
    s._size++
}


func (s *Stack) pop() rune {
    v := s._stack[s._size-1]
    s._size--
    return v
}

func (s *Stack) isEmpty() bool {
    return s._size == 0
}


/*
Valid Parentheses
====
Result: ACCEPTED
Runtime: 3 ms, faster than 22.13% of Go online submissions for Valid Parentheses.
Memory Usage: 2.1 MB, less than 49.77% of Go online submissions for Valid Parentheses.
=====
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
 

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.

*/