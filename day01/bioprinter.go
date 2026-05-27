package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	reader := bufio.NewReader(os.Stdin)


	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	fmt.Print("Enter your Age: ")
	strAge, _ := reader.ReadString('\n')
	strAge = strings.TrimSpace(strAge)

	//strconv.Atoi() : ASCII to Integer It converts: "21"  →  21 (string → int)

	
	age, err := strconv.Atoi(strAge)
	if err != nil {
		fmt.Println("Invalid age input. Please enter a number.")
		return
	}

	fmt.Print("Enter your City: ")
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	fmt.Print("Choose Go Level (beginner/intermediate/pro max): ")
	level, _ := reader.ReadString('\n')
	level = strings.TrimSpace(strings.ToLower(level))

	switch level {
	case "beginner":
		fmt.Println("You are starting your Go journey 🚀")

	case "intermediate":
		fmt.Println("You already know Go basics 💡")

	case "pro max":
		fmt.Println("You are a Go expert 🔥")

	default:
		fmt.Println("Invalid level selected")
		return
}

	birthYear := 2026 - age
	currentDay := 27
	daysLeft:=  365 - currentDay


	fmt.Println("\n╔══════════════════════════════════════╗")
    fmt.Println("║         GO DEVELOPER PROFILE         ║")
    fmt.Println("╠══════════════════════════════════════╣")

	fmt.Printf("║ Name       : %-24s║\n", name)	
	fmt.Printf("║ Age        : %-24d║\n", age)
	fmt.Printf("║ City       : %-24s║\n", city)
	fmt.Printf("║ Go Level   : %-24s║\n", level)
	fmt.Printf("║ Birth Year : %-24d║\n", birthYear)
	fmt.Printf("║ Days Left  : %-24d║\n", daysLeft)
	fmt.Println("╚══════════════════════════════════════╝")



	
}