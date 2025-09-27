package main

import "fmt"

func main() {
	var r float64
	const pi = 3.14

	fmt.Print("Masukkan jari-jari lingkaran : ")
	fmt.Scanln(&r)

	luas := pi * r * r
	fmt.Printf("Hasil dari penghitungan Luas lingkaran adalah %f ", luas)
}
