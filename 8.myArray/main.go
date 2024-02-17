package main

import "fmt"

func main() {
	fmt.Println(("welcome to array in go lang"))

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "mange"
	fruitList[3] = "peach"

	fmt.Println("Fruit liste is: ", fruitList)
	fmt.Println("Fruit liste is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushrom"}
	fmt.Println("Vegy list is: ", vegList)
}