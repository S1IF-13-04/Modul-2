package main

import "fmt"

func main() {
	const pi = 3.14159
	var jariJari, luas float64

	fmt.Print("Masukkan panjang jari-jari lingkaran: ")
	fmt.Scan(&jariJari)

	luas = pi * jariJari * jariJari

	fmt.Printf("Luas lingkaran dengan jari-jari %.1f adalah %.1f\n", jariJari, luas)
}