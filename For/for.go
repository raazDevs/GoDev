package main

import "fmt"

func main(){

	// i := 1
	// for i <= 10 {
	// 	println(i)
	// 	i = i + 4
	// }


	
	//infinite loop

	// for {
	// 	println("Hello World")
	// }




	//classic for loop

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }


	//Break and continue

	// for i := 1; i <= 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	if i == 8 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }	

	// Result: 1 2 3 4 6 7, because when i is 5, it will skip the print statement and continue to the next iteration, and when i is 8, it will break out of the loop.

	// Range loop: 1.55

	for i:= range 3 {
		fmt.Println(i)
	}



}

