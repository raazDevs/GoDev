package main

import "fmt"

// func main() {

// 	age := 20

// 	if age < 18 {
// 		fmt.Println("You are a minor.")
// 	} else if age >= 18 && age < 65 {
// 		fmt.Println("You are an adult.")
// 	} else {
// 		fmt.Println("You are a senior.")
// 	}
// }

//fallthrough — opt-in fall-through:


func main() {

	n := 2
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fallthrough   // explicitly continues to next case
	case 3:
		fmt.Println("three — printed because of fallthrough")
	case 4:
		fmt.Println("four — not printed, fallthrough stops")
   }

}



// main() calls add(10, 25)
//                 │
//                 ▼
// ┌─────────────────────────────────┐
// │  CALL STACK                     │
// │                                 │
// │  ┌─────────────────────────┐    │
// │  │  add()                  │    │
// │  │  a = 10  (copy of arg)  │    │
// │  │  b = 25  (copy of arg)  │    │
// │  │  result = 35            │    │
// │  │  return 35  ────────────┼─── ┼──→  sum = 35  (back in main)
// │  └─────────────────────────┘    │
// │                                 │
// │  ┌─────────────────────────┐    │
// │  │  main()                 │    │
// │  │  sum = ?  (waiting)     │    │
// │  └─────────────────────────┘    │
// └─────────────────────────────────┘

// Go passes arguments BY VALUE — a copy is made.
// Changing a inside add() does NOT change the original.
// (Pointers change this — we cover them on Day 5)

