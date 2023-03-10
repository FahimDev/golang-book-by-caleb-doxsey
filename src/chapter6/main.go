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

func assignMySlice() {
	/*
		A slice is a segment of an array. Like arrays slices are
		indexable and have a length. Unlike arrays this length
		is allowed to change. (Dynamic Array)
	*/
	var deviceCount int
	// We should use the built-in `make` function
	// 2nd parameter is fixing the current length of slice associated array
	// 3rd parameter is fixing the max capacity of slice associated array
	deviceSlice := make([]string, 0, 6)
	copySlice := make([]string, 5, 10)
	// ----------- R&D -----------
	fmt.Println("==>PHASE_1: Current Length of deviceSlice: ", len(deviceSlice))
	fmt.Printf("==>PHASE_1: Capacity of deviceSlice: %v\n\n", cap(deviceSlice))
	// ----------- R&D -----------
	defaultDevice := []string{"Buttscracher", "Plumbus"}
	fmt.Printf("Please, enter the number of your devices:")
	fmt.Scan(&deviceCount)
	// Go includes two built-in functions to assist with slices:
	// 1. append
	// 2. copy.
	deviceSlice = append(defaultDevice, "Golden Turd") // When we are Appending Slices the Capacity may get modified
	// ----------- R&D -----------
	fmt.Println("==>PHASE_2: Current Length of deviceSlice: ", len(deviceSlice))
	fmt.Println("==>PHASE_2: Slice: ", deviceSlice)
	fmt.Printf("==>PHASE_2: Capacity of deviceSlice: %v\n\n", cap(deviceSlice))
	// ----------- R&D -----------
	// OUTPUT: Buttscracher, Plumbus, Golden Turd
	for i := 0; i < deviceCount; i++ {
		var deviceName string
		fmt.Printf("Please, enter the Name of your device %v:", i+1)
		fmt.Scan(&deviceName)
		deviceSlice = append(deviceSlice, deviceName)
		// ----------- R&D -----------
		fmt.Println("==>PHASE_3: Current Length of deviceSlice: ", len(deviceSlice))
		fmt.Println("==>PHASE_3: Slice: ", deviceSlice) // When we are Appending Slices the Capacity may get modified
		fmt.Printf("==>PHASE_3: Capacity of deviceSlice: %v\n\n", cap(deviceSlice))
		// ----------- R&D -----------
	}
	// OUTPUT: // Buttscracher, Plumbus, Golden Turd, N number of input Values .....
	fmt.Println(deviceSlice)
	/*  The contents of deviceSlice are copied into copySlice. */
	copy(copySlice, deviceSlice)
	// copySlice may not have all the items copied. Because, we have defined its max length into 5.
	fmt.Printf("Updated copySlice: %v\n", copySlice)
	fmt.Println("==>PHASE_4: Current Length of copySlice: ", len(copySlice))
	fmt.Printf("==>PHASE_4: Capacity of copySlice: %v\n\n", cap(copySlice))
}

func main() {
	// showMyArray()
	// assignMyArray()
	assignMySlice()
}
