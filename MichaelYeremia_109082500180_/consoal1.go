package main

import "fmt"

func main() {
    var a, b, c, d, e int64 

    fmt.Println("Masukkan 5 bilangan bulat:")
    fmt.Scan(&a, &b, &c, &d, &e)

    hasil := a + b + c + d + e

    fmt.Println("Hasil penjumlahan =", hasil)
}
