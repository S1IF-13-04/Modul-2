package main

import "fmt"

func main() {
    var (
        satu, dua, tiga string
        temp            string
    )

    fmt.Print("Masukkan input String: ")
    fmt.Scan(&satu)
    fmt.Print("Masukkan Input Stirng: ")
    fmt.Scan(&dua)
    fmt.Print("Masukkan Input String: ")
    fmt.Scan(&tiga)
    fmt.Println("output awal = " + satu + " " + dua + " " + tiga)
    temp = satu
    satu = dua
    dua = tiga
    tiga = temp

    fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}