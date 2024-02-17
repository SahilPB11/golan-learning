package main

import "fmt"

const PublicVariable int = 250

func main() {
	var username string = "sahil"
	fmt.Println(username)
	fmt.Printf("Varibale is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Varibale is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Varibale is of type: %T \n", smallVal)

	var smallFloat float64 = 255.4545454454
	fmt.Println(smallFloat)
	fmt.Printf("Varibale is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable float64
	fmt.Println(anotherVariable)
	fmt.Printf("Varibale is of type: %T \n", anotherVariable)

	// implicit type
	var website = "hiii"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000.00
	fmt.Println(numberOfUser)


	// use public variable
	fmt.Println(PublicVariable)
	fmt.Printf("Varibale is of type: %T \n", PublicVariable)

}
