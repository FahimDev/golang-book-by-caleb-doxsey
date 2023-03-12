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

func makeClosure() func() uint {
	/* This method is returning a closure function which will return an uint value. */
	i := uint(0) // Declaring a variable with value 0.
	return func() (ret uint) {
		/*  A closure is a function value that references variables from outside its body. */
		ret = i // closure function setting return value as `i`.
		i += 2  // after setting the return value closure function is updating the value.
		return  // Defining the return syntax which will return the `rat` as non-updated `i` value.
	}
}

func main() {
	x, y := multiReturn()
	fmt.Printf("Return value 1: %v and return value 2: %v \n", x, y)
	fmt.Println("Total sum of parameters: ", variadicFunctions(1, 2, 3))

	/* Closure Test 1 */
	c := 0
	increment := func() int {
		c++
		return c
	}
	fmt.Println("Closure Test1: ", increment())
	fmt.Println("Closure Test1: ", increment())

	/* Closure Test 2 */
	// fmt.Println(makeClosure()) // Will print ==> 0x93b200 as closure function address
	nextEven := makeClosure()
	fmt.Println("Closure Test2: ", nextEven()) // 0
	fmt.Println("Closure Test2: ", nextEven()) // 2
	fmt.Println("Closure Test2: ", nextEven()) // 4
}
