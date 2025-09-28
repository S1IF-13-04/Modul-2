package main

import "fmt"

func main() {
	var nama, nim, kelas string

	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan Nim: ")
	fmt.Scanln(&nim)

	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&kelas)

	fmt.Println("Perkenalkan nama saya adalah", nama, ",salah satu mahasiswa Prodi S1-IF dari kelas", kelas, "dengan NIM,", nim)

}
