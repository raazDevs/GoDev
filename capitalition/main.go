package main

import "fmt"

func main() {
	var Public = "Data is  important"
	var private = "Data is  important"

	println(Public)
	println(private)

	// var price float32 = 50.343
	// fmt.Println(price)


	//shortend


	price:= 50.343
	fmt.Println(price)

}

//Capitalization in Go is used to determine the visibility of variables, functions, and other identifiers. If an identifier starts with an uppercase letter, it is exported and can be accessed from other packages. If it starts with a lowercase letter, it is unexported and can only be accessed within the same package. In the example above, "Public" is an exported variable, while "private" is unexported.
