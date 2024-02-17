package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var ptr2 = &myNumber
	fmt.Println("value of actual pointer is ", ptr2)
	fmt.Println("value of actual pointer is ", *ptr2)
	
	*ptr2 = *ptr2 * 2
	fmt.Println("value of actual pointer is ", myNumber)
	
}