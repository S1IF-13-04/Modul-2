package main

import "fmt"

func main() {
	var x int64
	fmt.Print("Masukan nilai x:")
	fmt.Scan(&x)
	fx := 2.0/(float64(x)+5) + 5
	fmt.Println(fx)
}

