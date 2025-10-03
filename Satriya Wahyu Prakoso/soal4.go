package main

import "fmt"

func main() {
	var F int
	var C float64
	fmt.Print("Masukkan Suhu dalam Fanrenheit: ")
	fmt.Scan(&F)
	C = float64(F-32) * 9.0 / 5.0
	fmt.Printf("Suhu dalam Celsius: %.2f\n", C)
}