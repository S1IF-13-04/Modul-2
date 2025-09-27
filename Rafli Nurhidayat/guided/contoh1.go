package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int

	fmt.Println("Masukkan 5 bilangan bulat:")
	fmt.Scan(&a, &b, &c, &d, &e)

	hasil := a + b + c + d + e

	fmt.Printf("Hasil penjumlahan: %d\n", hasil)
}
