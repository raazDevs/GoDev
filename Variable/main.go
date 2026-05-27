package main

import "fmt"

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

	// Short form
	x, y := 10, 20

	// Swap values (Go does this elegantly)
	x, y = y, x // x is now 20, y is now 10

	// Block declaration (great for related variables)
	// var (
	// 	host  = "localhost"
	// 	port  = 8080
	// 	debug = false
	// )

	// fmt.Println(host, port, debug)

	// var name string  = "Alice"
	// var age  int     = 25
	// var gpa  float64 = 8.5
	// var active bool  = true

	// fmt.Printf("Name: %s\n", name)      // %s = string
	// fmt.Printf("Age: %d\n", age)		// %d = integer (decimal)	
	// fmt.Printf("GPA: %.2f\n", gpa)	// %.2f = float, 2 decimal places
	// fmt.Printf("Active: %t\n", active) // %t = boolean (true/false)


	//Way 2: Short declaration with :=

	// name := "Saju"
	// age := 23
	// gpa := 8.75
	// active := true	
	// fmt.Printf("Name: %s\n", name)      // %s = string
	// fmt.Printf("Age: %d\n", age)		// %d = integer (decimal)	
	// fmt.Printf("GPA: %.2f\n", gpa)	// %.2f = float, 2 decimal places
	// fmt.Printf("Active: %t\n", active) // %t = boolean (true/false)	

	var price   float32 = 9.99       // 32-bit, less precise, less memory
	var salary  float64 = 95000.50   // 64-bit, more precise, recommended

	// When you write a decimal number, Go defaults to float64
	discount := 0.15    // float64 automatically
	pi       := 3.14159 // float64 automatically

	// Printf formatting
	fmt.Printf("%f\n", price)    // default float format: 9.990000
	fmt.Printf("%.2f\n", salary)   // 2 decimal places: 95000.50
	fmt.Printf("%f\n", discount)  // default float format: 0.150000
	
	fmt.Printf("%.4f\n", pi)       // 4 decimal places: 3.1416
	fmt.Printf("%e\n", salary)     // scientific: 9.500050e+04
}

//Rule: Use float64 by default. Only use float32 if you have a specific memory constraint.


// Type        Zero Value    Why it matters
//─────────────────────────────────────────────────
// int         0             Safe to do math with
// float64     0.0           Safe to do math with
// string      ""            Safe to compare and print
// bool        false         Safe for conditions
// pointer     nil           Safe — means "nothing here"
// slice       nil           Safe — means "empty list"
// map         nil           Safe — means "empty map"