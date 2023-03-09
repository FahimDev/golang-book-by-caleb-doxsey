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

	for i := 0; i < 3; i++ {
		fmt.Println(countryList[i])
	}
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
