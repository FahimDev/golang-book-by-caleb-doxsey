package main

import "fmt"

// escape sequences and comment

/*
Author: Md. Ariful Islam
Date: 01/03/2023
*/

func main() {

	testBool := true             // Boolean
	totalCGPA := 4               // Integer
	myCGPA := 3.26               // Floating point number
	myName := "Md. Ariful Islam" // String
	const COUNTRY = "Bangladesh"

	fmt.Printf("%v is working at \"Brain Station 23 Ltd.\"", myName)
	fmt.Printf("\nHe has got %v\\%v in BSc.\n", myCGPA, totalCGPA)
	fmt.Println("He have expertise in \n - Development and \n - Automation-QA \nboth side.")
	fmt.Println("He have published research \nin \t Blockchain.")
	fmt.Println("and \t IoT.")
	fmt.Printf("He is from %v \n", COUNTRY)
	fmt.Println(testBool)
}
