package main

import "fmt"

func main() {
	var nama string

	print("Masukan nama kamu:")
	fmt.Scanln(&nama)

	fmt.Println("Hallo", nama)

}
