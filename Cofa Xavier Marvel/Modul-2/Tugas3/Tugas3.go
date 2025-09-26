package main

import (
	"fmt"
	"math"
)

func main() {
	var JariJari float64
	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scanln(&JariJari)

	Luas := math.Pi * JariJari * JariJari
	fmt.Printf("Luas lingkaran dengan jari-jari %.f adalah %.1f\n", JariJari, Luas)
}
