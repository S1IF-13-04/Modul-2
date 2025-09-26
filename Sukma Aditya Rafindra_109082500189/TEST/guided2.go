package main

import "fmt"

func main() {
	var x int
	var fx float64

	fmt.Print("masukan nilai x: ")
	fmt.Scan(&x)
	fx = 2/float64(x+5) + 5
	fmt.Printf("nilai f(%d) = %f\n", x, fx)
}
