package main

import "fmt"

func main() {
	var fahrenheit float64
	var celsius float64
	fmt.Print("Enter temperature in Fahrenheit: ")
	fmt.Scanln(&fahrenheit)
	celsius = (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.f Fahrenheit is %.f Celsius\n", fahrenheit, celsius)
}
