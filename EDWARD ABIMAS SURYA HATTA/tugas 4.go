package main

import "fmt"

func main() {
    var f float64

    fmt.Println("masukkan suhu dalam fahrenheit : ")
    fmt.Scanln(&f)

    c := (f - 32) * 5 / 9

    fmt.Println("suhu dalam celcius adalah = ", c)
}
