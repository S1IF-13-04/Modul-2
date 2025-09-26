package main

import "fmt"

func main() {
    var nama, nim, kelas string

    fmt.Println("masukkan nama : ")
    fmt.Scanln(&nama)
    fmt.Println("masukkan nim : ")
    fmt.Scanln(&nim)
    fmt.Println("masukkan kelas : ")
    fmt.Scanln(&kelas)

    fmt.Printf("Perkenalkan saya adalah %s salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s\n", nama, kelas, nim)
}