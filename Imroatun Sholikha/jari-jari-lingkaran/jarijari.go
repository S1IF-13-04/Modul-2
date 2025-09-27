package main

import "fmt"

func main() {
	var pi float64 = 3.14159265359
	var jariJari float64
	fmt.Print("Masukkan panjang jari-jari lingkaran: ")
	fmt.Scanln(&jariJari)
	luas := pi * jariJari * jariJari
	fmt.Printf("Luas lingkaran adalah: %.1f", luas)
}
