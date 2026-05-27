package main

import "fmt"

func main(){
	fmt.Println("Hello, Go!")
	fmt.Print("Welcome to the world of Go programming.")
	fmt.Printf("Hello, %s\n", "Go")


	name := "Ravi"
	age  := 25
	gpa  := 8.75

	fmt.Printf("Name: %s\n", name)      // %s = string
	fmt.Printf("Age: %d\n", age)        // %d = integer (decimal)
	fmt.Printf("GPA: %.2f\n", gpa)      // %.2f = float, 2 decimal places
	fmt.Printf("Type: %T\n", name)       // %T = prints the type ("int")
	fmt.Printf("Value: %v\n", true)     // %v = any value, Go decides format
	}


//Critical rule in Go: if you import a package and never use it, your code will not compile.