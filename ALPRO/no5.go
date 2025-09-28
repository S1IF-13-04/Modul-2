package main

import "fmt"

func main() {
	var nama string
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&nama)

	fmt.Printf("Halo, nama saya adalah %s\n", nama)
}
