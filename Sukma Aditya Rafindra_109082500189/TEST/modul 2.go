package main

import "fmt"

func main() {
    var x int
	fmt.Print("masukkan bilangan bulat: ")
	fmt.Scan(&x)
	fx := 2.0 / (float64(x)5 + 5)
    
	// konversi ke float agar bisa menghitung pecahan
	fmt.Print ("hasil f(%d) = %v\n", x, fx)
}