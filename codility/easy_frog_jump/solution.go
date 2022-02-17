//https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp/
//Solution: OK 100%
package main

import (
	"fmt"
	"math"
)

func main() {
	r := Solution(10, 85, 30)
	fmt.Print(r)
}

func Solution(X int, Y int, D int) int {
	d := Y - X
	return int(math.Ceil(float64(d) / float64(D)))
}
