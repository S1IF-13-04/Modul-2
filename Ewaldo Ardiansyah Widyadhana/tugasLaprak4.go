package main

import (
	"fmt"
)

func main() {
	var fahrenheit int
	var celcius float64

	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	celcius = (5.0 / 9.0) * float64(fahrenheit-32)

	fmt.Printf("Suhu dalam Celcius: %.0f\n", celcius)
}
