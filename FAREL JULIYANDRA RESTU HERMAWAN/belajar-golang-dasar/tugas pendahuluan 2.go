package main

import "fmt"

func main() {
    var name string
    fmt.Print("Masukkan nama: ")
    fmt.Scan(&name)
    fmt.Println("Halo,", name)
}