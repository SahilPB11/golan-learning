package main

import (
	"fmt"
)

func main() {
	fmt.Print("dejfjj")

	var fruitList = []string{"Apple", "tomato", "peach"}

	fruitList = append(fruitList, "mango")
	fmt.Println("Type of fruitList is ", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 235
	highScore[2] = 236
	highScore[3] = 237

	highScore = append(highScore, 555, 458, 985)
	fmt.Println(highScore)

	// fmt.Println(highScore)
	// fmt.Println(sort.IntsAreSorted((highScore)))
	// sort.Ints(highScore)

	// how to remove value from based on index
	var courses = []string{"reactjs", "js", "swift", "python", "ruby"}
	fmt.Println(courses)
	
	var index int = 2
	
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
