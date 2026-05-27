package main

import "fmt"

func main(){
	a := 42
	b := 3.14
	c := "hello"
	d := true
	e := 'Z'
	f := 100_000_000

	fmt.Printf("a = %v  type: %T\n", a, a)
	fmt.Printf("b = %v  type: %T\n", b, b)
	fmt.Printf("c = %v  type: %T\n", c, c)
	fmt.Printf("d = %v  type: %T\n", d, d)
	fmt.Printf("e = %v  type: %T\n", e, e)
	fmt.Printf("f = %v  type: %T\n", f, f)
}