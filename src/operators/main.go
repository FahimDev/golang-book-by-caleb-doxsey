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

func logicalOperator(numOne int, numTwo int) {
	fmt.Printf("Is 1st input is not less than 2nd input: %v \n", !(numOne < numTwo))
	newVal := numOne > numTwo || numOne > 0
	fmt.Printf("Testing OR logical operator: %v \n", newVal)
	newVal = numOne > numTwo && numTwo > 10
	fmt.Printf("Testing AND logical operator: %v \n", newVal)
}

func bitwiseOperator(numOne int, numTwo int) {
	fmt.Printf("\n==> In Bitwise Operator Numbers will be calculated as Binary <==\n")
	fmt.Printf("& --> AND\n | --> OR\n ^ --> XOR\n")     //  << , >> are pending
	and := numOne & numTwo                               // if num1 = 18 (10010) and num2 = 17 (10001)
	fmt.Printf("%v AND %v = %v \n", numOne, numTwo, and) // then AND will be 16 (10000)
	or := numOne | numTwo
	fmt.Printf("%v OR %v = %v \n", numOne, numTwo, or)
	xor := numOne ^ numTwo
	fmt.Printf("%v XOR %v = %v \n", numOne, numTwo, xor)
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
	bitwiseOperator(numOne, numTwo)
	logicalOperator(numOne, numTwo)
}
