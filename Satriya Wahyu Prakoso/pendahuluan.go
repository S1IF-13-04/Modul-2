package main

import "fmt"

func main() {
	var nama string
	fmt.Print("Masukan nama anda: ")
	fmt.Scan(&nama)
	fmt.Print("Nama: ", nama)
}