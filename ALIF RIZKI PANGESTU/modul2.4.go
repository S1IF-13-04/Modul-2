package main

import (
	"fmt"
)

func main() {
	var fahrenheit int
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scan(&fahrenheit)
	celcius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("Suhu dalam Celcius: %d\n", celcius)
}
