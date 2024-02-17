package main

import "fmt"

func main() {
	fmt.Println("methods")

	// no inheritance in go lang
	// no super or parent

	hitesh := User{"sahil", "garg07825", true, 16}

	fmt.Println(hitesh)
	fmt.Printf("hites details are: %+v\n", hitesh)
	fmt.Printf("name is %v", hitesh.Name)
	hitesh.GetStatus()
	hitesh.NewMail("meow.com")
	fmt.Printf("name is %v, email is %v", hitesh.Name, hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(val string) {
	u.Email = val
	fmt.Println("Email of this user is: ", u.Email)
}
