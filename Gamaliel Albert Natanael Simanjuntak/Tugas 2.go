package main

import "fmt"

func main4() {
	var nama, nim, kelas string

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)

	fmt.Print("Masukkan kelas: ")
	fmt.Scan(&kelas)

	fmt.Println("Perkenalkan saya adalah", nama+", salah satu mahasiswa Prodi S1-IF dari kelas", kelas, "dengan NIM", nim+".")
}
