package main

import "fmt"

func main() {
    var r float64
    const pi = 3.14

    fmt.Println("Masukkan jari-jari lingkaran:")
    fmt.Scanln(&r)
    
    luas := pi * r * r
    
    fmt.Println("luas lingkaran adalah = ", luas)
}