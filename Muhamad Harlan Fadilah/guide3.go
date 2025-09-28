package main

import "fmt"

func main() {
	var c1, c2, c3, c4, c5 byte
	var b1, b2, b3 byte

	// input angka ASCII
	fmt.Scan(&c1, &c2, &c3, &c4, &c5)

	// input string (misalnya "SNO")
	var str string
	fmt.Scan(&str)

	// ambil 3 karakter dari string
	b1 = str[0]
	b2 = str[1]
	b3 = str[2]

	// output hasil
	fmt.Printf("%c%c%c%c%c\n", c1, c2, c3, c4, c5)
	fmt.Printf("%c%c%c\n", b1+1, b2+1, b3+1)
}
