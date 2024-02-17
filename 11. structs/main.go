package main

import "fmt"

func main() {
	fmt.Println("structs in go lang")

	// no inheritance in go lang
	// no super or parent

	hitesh := User{"sahil", "garg07825", true, 16}

	fmt.Println(hitesh)
	fmt.Printf("hites details are: %+v\n", hitesh)
	fmt.Printf("name is %v", hitesh.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
