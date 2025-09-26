package main

import "fmt"

func main() {
	var Fahrenheit, Celcius float64
	fmt.Printf("Fahrenheit :")
	fmt.Scan (&Fahrenheit)
	Celcius = ((Fahrenheit - 32) * 5 / 9)
	fmt.Println ("hasil", Celcius)
}