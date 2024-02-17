package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("world2")
	defer fmt.Println("world1")
	fmt.Println("defer")
	defer myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
