package main


import "fmt"


// func greet() {
// 	fmt.Println("Hello, From a Function!")
// }

// func main() {
// 	greet()
// 	greet()
// 	greet()
// }
	

// Function With Parameters


// func greetUser(name string) {
// 	fmt.Printf("Hello, %s! Welcome to Go programming.\n", name)
// }

// func add(a int, b int) {
// 	result := a + b
// 	fmt.Printf("The sum of %d and %d is %d.\n", a, b, result)
// }


// func multiply(a, b int) {
//     fmt.Println(a * b)
// }


// func main() {
// 	greetUser("Saju")
// 	greetUser("Rajdeep")
// 	add(5, 3)
// 	multiply(4, 6)
// }

//Function with a return value



// func add(a, b int) int {
// 	return a + b
// }

// func square(n float64) float64 {
// 	return n * n
// }

// func isEven(num int) bool {
// 	return num%2 == 0
// }


// func main() {	
// 	sum := add(5, 3)
// 	fmt.Printf("The sum: 5 + 3 = %d\n", sum)

// 	result := square(4.5)
// 	fmt.Printf("The square of 4.5 is %.2f\n", result)

// 	even := isEven(4)
// 	fmt.Printf("Is 4 even? %t\n", even)	

// 	fmt.Println(isEven(6))
// 	fmt.Println(isEven(7))

// }


//  Form 4 — Multiple return values (a Go superpower)



// func divide(a, b float64) (float64, error){
// 	if b == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return a / b, nil
// }


// func main(){
// 	result, err := divide(10, 2)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Printf("10/2 is %.2f\n", result)
// 		return
// 	}

// 	result, err = divide(10, 0)	
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Printf("10/0 is %.2f\n", result)
// 		return
// 	}	


//}


//-----------------------or-----------------------//


// func divide(a, b float64) (float64, error) {
//     if b == 0 {
//         return 0, fmt.Errorf("cannot divide by zero")
//     }
//     return a / b, nil   // nil means "no error"
// }

// func main() {
//     result, err := divide(10, 2)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     fmt.Printf("10 / 2 = %.2f\n", result)   // 10 / 2 = 5.00

//     result, err = divide(10, 0)
//     if err != nil {
//         fmt.Println("Error:", err)           // Error: cannot divide by zero
//         return
//     }
// }



//-------Named return values------//

// func minMax(nums []int)(min int, max int){
// 	min = nums[0]
// 	max = nums[0]

// 	for _, n := range nums {
// 		if n < min {
// 			min = n
// 		}
// 		if n > max {
// 			max = n
// 		}
// 	}
// 	return
// }

// func main(){
// 	low, high := minMax([]int{5, 2, 8, 1, 9})

// 	fmt.Printf("Low: %d, High: %d\n", low, high)
// }	


//--------Variadic functions---------//



// func sum(nums ...int) int {
//     total := 0
//     for _, n := range nums {
//         total += n
//     }
//     return total
// }

// func main() {
//     fmt.Println(sum(1, 2, 3))           // 6
//     fmt.Println(sum(10, 20, 30, 40))    // 100
//     fmt.Println(sum())                   // 0

//     // You can also pass a slice using ...
//     scores := []int{85, 90, 78, 92}
//     fmt.Println(sum(scores...))          // 345
// }



//----------------check the Average of numbers------------------//

// func average(nums ...int) float64 {
// 	if len(nums) == 0 {
// 		return 0
// 	}	
// 	total := 0
// 	for _, n := range nums {
// 		total += n
// 	}	
// 	return float64(total) / float64(len(nums))
// }
	
// func main() {
//     fmt.Println(average(1, 2, 3))           // 2.0
//     fmt.Println(average(10, 20, 30, 40))    // 25.0
//     fmt.Println(average())                   // 0.0
// }



// Accept zero or more ints
// func sum(nums ...int) int {
//     total := 0
//     for _, n := range nums {
//         total += n
//     }
//     return total
// }

// // Mixed: fixed params first, variadic last
// func logMessage(level string, parts ...string) {
//     fmt.Printf("[%s]", level)
//     for _, p := range parts {
//         fmt.Printf(" %s", p)
//     }
//     fmt.Println()
// }

// // Variadic with any type using interface{}
// func printAll(values ...interface{}) {
//     for i, v := range values {
//         fmt.Printf("%d: %v (%T)\n", i, v, v)
//     }
// }

// func main() {
//     fmt.Println(sum())              // 0   — zero args is fine
//     fmt.Println(sum(1, 2, 3))      // 6
//     fmt.Println(sum(10, 20, 30, 40, 50))  // 150

//     // Pass a slice by spreading it with ...
//     scores := []int{85, 92, 78, 96}
//     fmt.Println(sum(scores...))    // 351

//     logMessage("INFO", "server", "started", "on", "port", "8080")
//     // [INFO] server started on port 8080

//     printAll("hello", 42, true, 3.14)
//     // 0: hello (string)
//     // 1: 42 (int)
//     // 2: true (bool)
//     // 3: 3.14 (float64)
// }



//-------Form 7 — Anonymous functions and closures----//\



// func main() {
//     // Define and call immediately
//     func() {
//         fmt.Println("I am anonymous!")
//     }()    // the () at the end calls it right away

//     // Assign to a variable
//     double := func(n int) int {
//         return n * 2
//     }
//     fmt.Println(double(7))   // 14

//     // Closure — a function that "captures" a variable from its outer scope
//     counter := 0
//     increment := func() {
//         counter++   // captures and modifies counter from outer scope
//     }
//     increment()
//     increment()
//     increment()
//     fmt.Println(counter)   // 3
// }



func main() {
    // Define and call immediately
    func() {
        fmt.Println("I am anonymous!")
    }()    // the () at the end calls it right away

    // Assign to a variable
    double := func(n int) int {
        return n * 2
    }
    fmt.Println(double(7))   // 14

    // Closure — a function that "captures" a variable from its outer scope
    counter := 0
    increment := func() {
        counter++   // captures and modifies counter from outer scope
    }
    increment()
    increment()
    increment()
    fmt.Println(counter)   // 3
}






func readFile() {
    fmt.Println("1. Opening file")
    defer fmt.Println("4. Closing file")   // runs LAST

    fmt.Println("2. Reading file")
    fmt.Println("3. Processing data")
    // function ends here, then deferred call runs
}

func main() {
    readFile()
}



func main() {
    defer fmt.Println("first defer — runs LAST")
    defer fmt.Println("second defer — runs MIDDLE")
    defer fmt.Println("third defer — runs FIRST")
    fmt.Println("main body")
}
// Output:
// main body
// third defer — runs FIRST
// second defer — runs MIDDLE
// first defer — runs LAST


