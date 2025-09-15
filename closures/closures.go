package main

import (
	"fmt"
	""
)

// makeCounter returns a closure
func makeCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor // captures 'factor'
	}
}

func main() {
	counter := 0

	
// 👉 Here:
// increment is a closure.
// It captures the variable counter from main.
// Even though main is running, the closure keeps its own access to counter.

	// this is a closure
	increment := func() int {
		counter++          // captures & modifies counter
		return counter
	}

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3

//Ex:2 Multiple Closures with Independent States
	c1 := makeCounter()
	c2 := makeCounter()

	fmt.Println(c1()) // 1
	fmt.Println(c1()) // 2
	fmt.Println(c2()) // 1 (different closure, new counter)
	fmt.Println(c1()) // 3


//Ex:3 Closure Capturing Variables from Enclosing Scope
	double := multiplier(2)
	triple := multiplier(3)
// double and triple both are closures that "remember" the factor they were created with.
	fmt.Println(double(5)) // 10
	fmt.Println(triple(5)) // 15
	fmt.Println(triple(5)) // 15


}

