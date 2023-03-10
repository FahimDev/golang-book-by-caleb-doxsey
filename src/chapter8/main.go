package main

// Pointers | Call by Reference

import "fmt"

/*
Author: Md. Ariful Islam
Date: 01/03/2023
*/

func changeValue(val int) {
	val = val - 86
}

func callByReference(myPointer *int) {
	fmt.Println("Messing at direct memory location...")
	*myPointer = *myPointer - 86
}

func testingNewFunc(xPtr *int) {
	*xPtr = 7286837 // Assigning Value
}

type student struct {
	Name string
	Roll int
}

func experimentObj() *student {
	structInstance := student{Name: "Fahim", Roll: 25}
	fmt.Printf("Student Obj: %p \n", &structInstance)
	fmt.Printf("Address of struct = %+v: %p\n", structInstance, &structInstance)
	fmt.Printf("Address of struct field = %s: %p\n", structInstance.Name, &structInstance.Name)
	fmt.Printf("Address of struct field = %d: %p\n", structInstance.Roll, &structInstance.Roll)
	return &structInstance
}

func main() {
	var point *int // This can store address
	x := 786
	fmt.Println(x)
	point = &x
	fmt.Println(&x)
	fmt.Println(point)  // This will print the allocated address
	fmt.Println(*point) // This will print the stored value in the allocated address

	changeValue(x)
	fmt.Println("Called changeValue() but no change in value =>", x)
	// Messing at direct memory location...
	callByReference(&x)
	fmt.Println(*point)
	fmt.Println(x) // The main value also got updated.

	// Another way to get a pointer is to use the built-in new function
	xPtr := new(int)
	testingNewFunc(xPtr)
	fmt.Println("Testing New method value: ", *xPtr)
	/*
	 Go is a garbage collected programming language
	 which means memory is cleaned up automatically
	 when nothing refers to it anymore.
	*/
	fmt.Printf("MAIN ==> %p\n", experimentObj())

}
