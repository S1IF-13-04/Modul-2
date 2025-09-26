package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var hasil int
	fmt.Println("Masukkan 5 angka yang ingin dijumlahkan (pisahkan dengan spasi):")
	fmt.Scanln(&a, &b, &c, &d, &e)
	hasil = a + b + c + d + e
	fmt.Println("Hasil penjumlahan:", hasil)

}
