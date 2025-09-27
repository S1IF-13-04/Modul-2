package main

import "fmt"

func main() {
	var x, fx float64

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fx = 2/(x+5) + 5

	fmt.Printf("Hasil f(%.1f) = %.17g\n", x, fx)
}