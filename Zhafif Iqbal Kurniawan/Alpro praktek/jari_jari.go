package main

import (
	"fmt"
	"math"
)

func main() {

	var r float64

	fmt.Print("Masukkan nilai jari jari lingkaran :")
	fmt.Scanln(&r)

	l := math.Pi * r * r

	fmt.Println("Hasil :")
	fmt.Println("Luas :", l)

}
