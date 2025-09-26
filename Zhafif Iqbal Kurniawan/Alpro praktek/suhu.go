package main

import (
	"fmt"
)

func main() {

	var a float64

	fmt.Print("Masukkan suhu fahrenheit :")
	fmt.Scanln(&a)

	c := (a - 32) * 5 / 9

	fmt.Println("Celcius :", c)

}
