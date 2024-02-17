package main

import (
	"fmt"

)

func main() {
	fmt.Println("If Else")

	loginCount := 43

	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "something else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10{
		fmt.Println("Num is less then 10")
	}else{
		fmt.Println("Num is not less then 10")
	}

}
