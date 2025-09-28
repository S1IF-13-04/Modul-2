package main

import "fmt"

func main() {
	var Fahrenheit int
	var Celcius int

	fmt.Print("Masukkan suhu Fahrenheit: ")
	fmt.Scanln(&Fahrenheit)


	Celcius = (Fahrenheit - 32) * 5 / 9

	fmt.Println("Suhu dalam Celcius =", Celcius)
}
