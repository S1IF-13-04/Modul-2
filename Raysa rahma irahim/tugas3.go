package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Print("Masukkan jari-jari: ")
	fmt.Scan(&r)

	luas := math.Pi * r * r
	fmt.Printf("%.1f\n", luas)
}
