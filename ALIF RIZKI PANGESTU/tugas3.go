package main

import "fmt"

func main() {
	// input 5 bilangan integer
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	// cetak dalam format karakter (tidak ada spasi)
	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	// input 3 karakter
	var x, y, z byte
	fmt.Scanf("%c%c%c", &x, &y, &z)

	// cetak 3 karakter berikutnya dalam tabel ASCII
	fmt.Printf("%c%c%c\n", x+1, y+1, z+1)
}
