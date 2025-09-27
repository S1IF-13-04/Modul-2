package main

import "fmt"

func main() {
	var F float64
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&F)
	C := (F - 32) * 5 / 9
	fmt.Println("Suhu dalam Celsius =", C)
}
