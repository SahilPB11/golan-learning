package main

import "fmt"

func main() {
	fmt.Println("Functions ")
	greeter()

	result := adder(3, 4)
	fmt.Println("Resukt is : ", result)
	
	result1 := proAder(5, 7, 8, 9, 4, 5, 6, 8)
	fmt.Println("Resukt is : ", result1)
}

func greeter() {
	fmt.Println("greeter 1")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
