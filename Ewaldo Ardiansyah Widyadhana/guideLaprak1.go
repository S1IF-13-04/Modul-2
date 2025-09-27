package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var hasil int
	fmt.Scanln(&a, &b, &c, &d, &e)
	hasil = a + b + c + d + e
	fmt.println("hasil penjumlahan", a, b, c, d, e, "adalah", hasil)

}
