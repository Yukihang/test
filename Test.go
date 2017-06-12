package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func run( a int, b int) int {
	return a + b
}

func main() {
	var a = 1
	fmt.Println(a)
	fmt.Println(run(1,2))
}
