package main

import (
	"fmt"
)

func main() {
	var fahrenheit int
	fmt.Print("Masukan suhu dalam Fahrenheit: ")
	fmt.Scan(&fahrenheit)
	celcius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("suhu dalam Celcius) %d\n", celcius)
}
