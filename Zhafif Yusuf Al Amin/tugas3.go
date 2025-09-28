package main

import "fmt"

func main() {
	var r float64
	const pi = 3.14159
	var luas float64

	fmt.Print("Masukkan jari jari lingkaran: ")
	fmt.Scanln(&r)

	luas = pi * r * r
	fmt.Printf("Luas lingkaran adalah = %.1f\n", luas)
}
