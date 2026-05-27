package main

import "fmt"

func main() {
	name := "Rajdeep"
	city := "Morigaon"
	empty := ""

	greeting := "Hello, " + name + "! I live in " + city + empty
	fmt.Println(greeting)

	fmt.Println(len("Hello"))   
	fmt.Println(len("नमस्ते"))

	message := `Dear User,
	Welcome to Go programming.
	This string preserves all whitespace and newlines exactly.
	No need to escape anything here.`

	fmt.Println(message)

	result := fmt.Sprintf("Name:%s, Age:%d", name, 25)
	fmt.Println(result)
}