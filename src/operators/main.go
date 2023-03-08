package main

import "fmt"

/*
Author: Md. Ariful Islam
Date: 01/03/2023
*/

// Types of operators
// 1. Arithmetic Operator
// 2. Assignment Operator
// 3. Unary Operators
// 4. Relational Operators
// 5. Logical Operator
// 6. Bitwise Operator
// 7. Special Operator

func arithmeticOperator(numOne int, numTwo int) {
	var result = numOne + numTwo
	fmt.Printf("%v + %v = %v\n", numOne, numTwo, result)

	result = numOne - numTwo
	fmt.Printf("%v - %v = %v\n", numOne, numTwo, result)

	result = numOne * numTwo
	fmt.Printf("%v × %v = %v\n", numOne, numTwo, result)

	var resultFloat = float32(numOne) / float32(numTwo)
	fmt.Printf("%v ÷ %v = %v\n", numOne, numTwo, resultFloat)

	result = numOne % numTwo
	fmt.Printf("Remainder of (%v ÷ %v) is %v\n", numOne, numTwo, result)
}

func assignmentOperator(numOne int, numTwo int) {
	numOne += numTwo // numOne = numOne + numTwo
	fmt.Printf("Assignment Operator (+) Value: %v \n", numOne)
	numOne -= numTwo // numOne = numOne - numTwo
	fmt.Printf("Assignment Operator (-) Value: %v \n", numOne)
}

func unaryOperator(numOne int, numTwo int) {
	numOne++
	fmt.Printf("After Post Increment of 1st Input: %v \n", numOne)
	numTwo--
	fmt.Printf("After Post Decrement of 2st Input: %v \n", numTwo)
}

func relationalOperator(numOne int, numTwo int) {
	fmt.Printf("Is 1st input less than 2nd input: %v \n", numOne < numTwo)
	fmt.Printf("Is 1st input greater than 2nd input: %v \n", numOne > numTwo)
	fmt.Printf("Is 1st input equal to 2nd input: %v \n", numOne == numTwo)
	fmt.Printf("Is 1st input not-equal to 2nd input: %v \n", numOne != numTwo)
}

func main() {
	var numOne, numTwo int
	fmt.Printf("Input Number 1: ")
	fmt.Scan(&numOne)
	fmt.Printf("Input Number 2: ")
	fmt.Scan(&numTwo)

	relationalOperator(numOne, numTwo)
	arithmeticOperator(numOne, numTwo)
	assignmentOperator(numOne, numTwo)
	unaryOperator(numOne, numTwo)
}
