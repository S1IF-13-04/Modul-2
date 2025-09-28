package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64 // jari-jari lingkaran

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&r)

	// rumus luas = π * r * r
	luas := math.Pi * r * r

	fmt.Printf("Luas lingkaran = %.1f\n", luas)
}
