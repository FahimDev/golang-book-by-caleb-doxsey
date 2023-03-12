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

func factorialRecursion(x uint) uint {
	if x == 0 {
		fmt.Println("Bottom Found!")
		return 1
	}
	fmt.Println("Recursion....")
	return x * factorialRecursion(x-1)
}

func deferTest() {
	// defer works more like LIFO
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func deferTest1() (i int) {
	defer func() { i++ }()
	return 1
}

func recoverTest() {
	defer func() {
		if r := recover(); r != nil {
			// Recover function is getting the value whatever Panic function is returning...
			fmt.Println("Recovered in recoverTest", r)
		}
	}()
	fmt.Println("Calling panicTest.")
	panicTest(0)
	fmt.Println("Returned normally from panicTest.")
}

func panicTest(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v <#", i))
	}
	defer fmt.Println("Defer in panicTest", i) // LIFO
	fmt.Println("Printing in panicTest", i)
	panicTest(i + 1)
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

	fmt.Println("Recursion Test: ", factorialRecursion(5))

	fmt.Println("Defer Test Start")
	fmt.Println("Defer Test 0")
	deferTest()
	fmt.Println("Defer Test 1")
	fmt.Println("Defer: ", deferTest1())
	fmt.Println("Defer Test End")

	recoverTest() // https://go.dev/blog/defer-panic-and-recover
	fmt.Println("Returned normally from recoverTest.")

}
