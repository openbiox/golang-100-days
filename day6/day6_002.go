// Echo4 prints its command-line arguments.
package main

import (
	"fmt"
)

func main() {
	p := newInt()
	q := newInt()
	fmt.Println(p == q) // "false"

	p = newInt()
	q = newInt2()
	fmt.Println(p == q) // "false"
}

// newInt equal newInt2 func
func newInt() *int {
	return new(int)
}
func newInt2() *int {
	var dummy int
	return &dummy
}

// greatest common divisor
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

//Fibonacci

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
