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

	var domainOne string
	var domainTwo string
	var workPlace string

	fmt.Print("Enter to two domain names of your expertise: ")
	// Multiple Input at a time
	fmt.Scanf("%v %v", &domainOne, &domainTwo)

	fmt.Print("Enter the name of your workplace: ")
	// Single Input
	fmt.Scan(&workPlace)

	fmt.Printf("%v is working at \"%v\"", myName, workPlace)
	fmt.Printf("\nHe has got %v\\%v in BSc.\n", myCGPA, totalCGPA)
	fmt.Println("He have expertise in \n - Development and \n - Automation-QA \nboth side.")
	fmt.Printf("\nHe have published research \nin \t %v", domainOne)
	fmt.Printf("\nand \t %v.", domainTwo)
	fmt.Printf("\nHe is from %v \n", COUNTRY)
	fmt.Println(testBool)
}
