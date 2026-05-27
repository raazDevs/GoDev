package main

import "fmt"

func main() {
    tempC := 37
    tempF := (float64(tempC) * 9.0 / 5.0) + 32.0
    isNormal := tempF >= 97.0 && tempF <= 99.0

    fmt.Printf("%d°C = %.2f°F\n", tempC, tempF)
    fmt.Printf("Normal body temperature: %v\n", isNormal)
}