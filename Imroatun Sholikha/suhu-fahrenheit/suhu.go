package main

import "fmt"

func main() {

	var fahrenheit int

	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	celsius := (float64(fahrenheit) - 32) * 5 / 9
	fmt.Printf("Suhu dalam Celsius adalah: %.0f", celsius)

}
