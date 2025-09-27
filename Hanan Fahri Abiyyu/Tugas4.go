package main

import "fmt"

func main() {
	var f float64

	fmt.Print("Suhu Fahrenheit :")
	fmt.Scanln(&f)

	c := (f - 32) * 5 / 9

	fmt.Printf("Hasil suhu dalam Celcius = %f ", c)

}
