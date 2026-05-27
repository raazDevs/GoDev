package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	input, _ := reader.ReadString('\n')
	name := strings.TrimSpace(input) // Remove newline character
	fmt.Printf("Hello, %s!\nWelcome to Go programming!\n", name)
}

// reader := bufio.NewReader(os.Stdin)

//fmt.Print("Enter your name: ")
//input,_ := reader.ReadString('\n')
