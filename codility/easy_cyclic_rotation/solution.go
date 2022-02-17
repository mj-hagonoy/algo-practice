//https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/
//Solution OK 100%
package main

func main() {

}

func Solution(A []int, K int) []int {
	if len(A) == 0 || A == nil {
		return A
	}
	if K == len(A) {
		return A
	}

	for i := 0; i < K; i++ {
		last := len(A) - 1
		a1 := A[last:]
		b := append(a1, A[0:last]...)
		A = b
	}

	return A
}
