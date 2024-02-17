package main

import "fmt"

func main() {
	fmt.Println("Maps in go lang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["rc"] = "react"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts: ", languages["js"])

	delete(languages, "rb")
	fmt.Println("List of all languages: ", languages)

	// loops are intersting in go lang

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
