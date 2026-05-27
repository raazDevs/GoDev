package main

import "fmt"


func main(){
	isLoggedIn  := true
	isAdmin     := false
	

// Logical operators
fmt.Println(isLoggedIn && isAdmin)    // AND: both must be true → false
fmt.Println(isLoggedIn || isAdmin)    // OR:  at least one true → true
fmt.Println(!isLoggedIn)              // NOT: flips the value → false

// Comparison operators produce booleans
age := 20
fmt.Println(age >= 18)    // true
fmt.Println(age == 21)    // false
fmt.Println(age != 18)    // true
}