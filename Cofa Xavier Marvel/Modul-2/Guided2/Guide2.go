package main

import "fmt"

func main() {

	var X float64
	fmt.Print("Masukkan nilai X: ")
	fmt.Scan(&X)
	fx := 2/(X+5) + 5
	fmt.Printf("f(x) = 2/(%v + 5) + 5 = %v ", X, fx)
}
