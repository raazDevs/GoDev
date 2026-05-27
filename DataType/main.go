package main

import "fmt"

func main() {

	// Signed integers (can be negative)
	// var a int8 = 127                  // -128 to 127
	// var b int16 = 32767               // -32768 to 32767
	// var c int32 = 2147483647          // about -2 billion to 2 billion
	// var d int64 = 9223372036854775807 // very large

	// // Most common: just use int (64-bit on modern systems)
	// age := 25 // type is int

	// // Unsigned integers (can't be negative, larger positive range)
	// var e uint8 = 255 // 0 to 255
	// var f uint64 = 1000000

	// // Special aliases
	// var g byte = 65  // byte = uint8, used for raw data
	// var h rune = 'A' // rune = int32, used for Unicode characters

	name := "Ravi"
	age := 25
	gpa := 8.75

	fmt.Printf("Name: %s\n", name)  // %s = string
	fmt.Printf("Age: %d\n", age)    // %d = integer (decimal)
	fmt.Printf("GPA: %.2f\n", gpa)  // %.2f = float, 2 decimal places
	fmt.Printf("Type: %T\n", age)   // %T = prints the type ("int")
	fmt.Printf("Value: %v\n", true) // %v = any value, Go decides format

}