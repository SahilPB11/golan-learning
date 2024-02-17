package main

import "fmt"

func main() {
	fmt.Println("loops")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "friday", "saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	roudgeValue := 2

	for roudgeValue < 10 {

		if roudgeValue == 2 {
			goto lco
		}

		if roudgeValue == 5 {
			break
		}
		fmt.Println("Value is ", roudgeValue)
		roudgeValue++
	}

lco:
	fmt.Println("jumping at code is it ok")

}
