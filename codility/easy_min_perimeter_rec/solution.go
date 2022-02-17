package main

import "fmt"

func main(){
	fmt.Println(Solution(30)) //expected 22
	fmt.Println(Solution(1)) //expected 4
}

func Solution(N int) int {
	return getMinPerimeter(N)
 }
 
 func getMinPerimeter(N int) int {
	 if N == 1 {return 2 * (N+N)} //1*1 = 1, and 2 * (1+1) = 4
	 min := 2 * (1 + N)  // 1 and N
	 i := 2
	 for i * i < N {
		 if N % i == 0 {
			 p := 2 * ((N/i) + i)
			 if p < min {
				 min = p
			 }
		 }
		 i++
	 }
 
	 if i * i == N {
		 p := 2 * ((N/i) + i)
			 if p < min {
				 min = p
			 }
	 }
 
	 return min
 }

 /*
 TASK SCORE : 100%
 An integer N is given, representing the area of some rectangle.

The area of a rectangle whose sides are of length A and B is A * B, and the perimeter is 2 * (A + B).

The goal is to find the minimal perimeter of any rectangle whose area equals N. The sides of this rectangle should be only integers.

For example, given integer N = 30, rectangles of area 30 are:

(1, 30), with a perimeter of 62,
(2, 15), with a perimeter of 34,
(3, 10), with a perimeter of 26,
(5, 6), with a perimeter of 22.
Write a function:

func Solution(N int) int

that, given an integer N, returns the minimal perimeter of any rectangle whose area is exactly equal to N.

For example, given an integer N = 30, the function should return 22, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..1,000,000,000].
 
 */