package main
import "fmt"

const age = 24

func main(){
	// const name = "go lang"

	// name = "heavy dukh"

	fmt.Println(age)

	const(
		port = 8080
		host = "localhost"
	)
	fmt.Println(port, host)

}

//outer side of the main function, we can declare a constant using the const keyword. A constant is a value that cannot be changed once it is assigned. In the example above, we declare a constant named 'name' and assign it the value "go lang".


// in the case of shorthand syntax, we cannot use it to declare constants. The shorthand syntax (:=) is only used for variable declaration and initialization, not for constants. Constants must be declared using the const keyword and cannot be reassigned after their initial assignment.

// For example, the following code will result in a compilation error:
// name := "go lang" // This will cause an error because constants cannot be declared using shorthand syntax
// const name = "go lang" // This is the correct way to declare a constant

// for example, if we try to reassign a value to a constant, it will also result in a compilation error:
// const name = "go lang"
// name = "heavy dukh" // This will cause an error because constants cannot be reassigned








