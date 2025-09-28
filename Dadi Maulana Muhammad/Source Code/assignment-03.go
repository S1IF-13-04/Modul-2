//Algoritma
//Program suhu converter dari fahrenheit ke celcius
//Kamus : var F float64
//Input : celcius
//				celcius = (F - 32) * 5 / 9
//Output : celsius
//endProgram

package main

import (
	"fmt"
)

func main() {
	var F float64
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scan(&F)

	C := (F - 32) * 5 / 9
	fmt.Printf("%.0f\n", C)
}
