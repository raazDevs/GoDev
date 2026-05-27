// package main

// import (
//     "fmt"
//     "os"
// )

// func main() {
//     var username string
//     username := "devuser"
//     fmt.Println(username)
// }

// --------------- Debugging the code above -----------------//

//But := is only used for creating a new variable.

// package main

// import "fmt"

// func main() {
//     var username string
//     username = "devuser" // Assigning value to the already declared variable
//     fmt.Println(username)
// }

//------------------------2---------------------------

// package main
// import "fmt"

// func main() {
//     score := 87
//     percentage := 99.5
//     combined := score + percentage
//     fmt.Printf("Combined: %f\n", combined)
// }

// package main
// import "fmt"

// func main() {
//     score := 87
//     percentage := 99.5
//     combined := float64(score) + percentage // Convert score to float64 before adding
//     fmt.Printf("Combined: %f\n", combined)
// }

// ------------------3------------------------//

//fixed value that cannot be modified//
// package main
// import "fmt"

// const MaxConnections = 100

// func main() {
//     //MaxConnections = 200 // This will cause a compile-time error because MaxConnections is a constant
//     fmt.Println("Max:", MaxConnections)
// }

//---------------------4-------------------//

package main

import "fmt"

func main() {
    x, y := 5, 10
    //x, y := y, x 
    // This creates new variables x and y that shadow the outer ones, instead of swapping their values
    x , y = y, x // This correctly swaps the values of x and y
    fmt.Println(x, y)
}