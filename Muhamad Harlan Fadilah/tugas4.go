package main

import "fmt"

func main() {
	var fahrenheit float64
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	celcius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("Suhu dalam Celcius: %.2f", celcius)
}
