package main

/*
Functions
1. Returning Multiple Values
2. Variadic Functions
3. Closure
4. Recursion
5. Defer, Panic & Recover
*/

import "fmt"

func multiReturn() (int, int) {
	return 5, 6
}

func variadicFunctions(args ...int) int {
	/* There is a special form available for the last parameter in Go.
	By using `...` before the type name of the last parameter you can indicate
	that it takes zero or more of those parameters */
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	x, y := multiReturn()
	fmt.Printf("Return value 1: %v and return value 2: %v \n", x, y)
	fmt.Println("Total sum of parameters: ", variadicFunctions(1, 2, 3))
}
