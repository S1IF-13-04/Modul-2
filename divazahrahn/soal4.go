package main

import "fmt"

func main() {
	var fahrenheit int
	fmt.Print("Masukan suhu fahrenheit:")
	fmt.Scan(&fahrenheit)

	celcius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("suhu dalam celcius %d \n ", celcius)
}
