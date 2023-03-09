package main

import "fmt"

// control structures

/*
Author: Md. Ariful Islam
Date: 01/03/2023
*/

func main() {
	var inputNum int
	fmt.Printf("Please, entry an integer number: ")
	fmt.Scan(&inputNum)

	switch inputNum {
	case 0:
		fmt.Println("Zero for switch trial")
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		fmt.Println("The one-digit number are engaged")
	default:
		fmt.Println("Executing switch default...")
		if inputNum%2 != 0 {
			i := 1
			for i <= 10 {
				fmt.Println(i)
				i++
			}
		} else {
			for i := 10; i >= 1; i-- {
				fmt.Println(i)
			}
		}
	}

}
