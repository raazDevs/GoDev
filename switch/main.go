package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 2
	// switch i {
	// case 1:
	// 	println("one")
	// case 2:
	// 	println("two")
	// case 3:
	// 	println("three")
	// default:
	// 	println("other")
	// }

	//Multiple cases

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")	
	default:
		fmt.Println("It's a weekday")
	}
}