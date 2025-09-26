package main

import (
	"fmt"
)

func main() {

	var a, b, c string

	fmt.Print("Masukan nama Anda ")
	fmt.Scan(&a, &b, &c)
	fmt.Println("Halo", a, b, c)
}
