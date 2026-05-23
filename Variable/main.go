package main

func main() {
	// var name string = "golang"

	// infer
	// var name = "golang"
	// var isAdult bool = true

	// var age int = 30

	// shorthand syntax
	// name := "golang"

	// var name string
	// name = "golang"

	// var price float32 = 50.5
	// var price = 50.5
	// price := 50.5

	// fmt.Println(price)
}
















//the main thing is that may be you want to declare a variable without initializing it, in that case you can use the var keyword without assigning a value to it. This will give the variable a default value based on its type (e.g., 0 for int, "" for string, false for bool).

// For example:	Declaring a variable without initializing it
// var name string
// fmt.Println(name) // Output: ""		
// In this case, the variable 'name' is declared as a string but not initialized, so it will have the default value of an empty string ("").
// You can also declare multiple variables of the same type in a single line using the var keyword. For example:
// var age, height int
// fmt.Println(age, height) // Output: 0 0
// Here, both 'age' and 'height' are declared as integers and will have the default value of 0.
// Additionally, you can declare and initialize multiple variables in a single line using the shorthand syntax. For example:
// name, age := "golang", 30
// fmt.Println(name, age) // Output: golang 30
// In this case, 'name' is initialized with the value "golang" and 'age' is initialized with the value 30 using the shorthand syntax.
// Overall, the var keyword is used to declare variables in Go, and it can be used with or without initialization. The shorthand syntax provides a convenient way to declare and initialize variables in a single line.
// In Go, you can declare variables using the var keyword. The var keyword allows you to specify the type of the variable and optionally initialize it with a value. Here are some examples of how to declare variables in Go:

// Declaring a variable with a specific type and initializing it
// var name string = "golang"
// Declaring a variable with type inference (the type is inferred from the assigned value)
// var name = "golang"
// Declaring a variable using shorthand syntax (only for local variables)

// name := "golang"
// Declaring a variable without initializing it (it will have a default value based on its type)
// var name string
// fmt.Println(name) // Output: ""	

