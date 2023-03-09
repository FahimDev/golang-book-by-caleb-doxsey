package main

// Arrays, Slices and Maps

import "fmt"

/*
Author: Md. Ariful Islam
Date: 01/03/2023
*/

func showMyArray() {
	var countryList = [3]string{
		"Bangladesh",
		"India",
		"Japan",
	}
	// The Go compiler won't allow you to create variables that you never use.
	// A single _ (underscore) is used to tell the compiler that we don't need this.
	for _, value := range countryList {
		fmt.Println(value)
	}
	fmt.Printf("The array length is: %v", len(countryList))
}

func assignMyArray() {
	var myArray [5]string

	for i := 0; i < 5; i++ {
		fmt.Printf("Input an array item and press Enter:")
		fmt.Scan(&myArray[i])
	}

	fmt.Println(myArray)
}

func main() {
	showMyArray()
	assignMyArray()
}
