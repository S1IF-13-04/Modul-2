package main

import "fmt"

func main() {
	var fahrenheit, celsius float64

	// Input suhu Fahrenheit
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	// Konversi: C = (F - 32) × 5/9
	celsius = (fahrenheit - 32) * 5 / 9

	// Output hasil konversi
	fmt.Printf("%.0f°F = %.0f°C\n", fahrenheit, celsius)
}