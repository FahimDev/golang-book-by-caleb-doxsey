package main

import "fmt"

// escape sequences and comment

/*
Author: Md. Ariful Islam
Date: 01/03/2023
*/

func main() {

	var testBool bool = true               // Boolean
	var totalCGPA int = 4                  // Integer
	var myCGPA float32 = 3.26              // Floating point number
	var myName string = "Md. Ariful Islam" // String

	fmt.Printf("%s is working at \"Brain Station 23 Ltd.\"", myName)
	fmt.Printf("\nHis has got %.2f\\%d in BSc.\n", myCGPA, totalCGPA)
	fmt.Println("He have expertise in \n - Development and \n - Automation-QA \nboth side.")
	fmt.Println("He have published research \nin \t Blockchain.")
	fmt.Println("and \t IoT.")
	fmt.Println(testBool)
}
