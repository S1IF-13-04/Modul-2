package main

import "fmt"

func main() {
	var nama, nim, kelas string
	fmt.Print("Masukkan nama   : ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan NIM    : ")
	fmt.Scanln(&nim)
	fmt.Print("Masukkan Kelas  : ")
	fmt.Scanln(&kelas)
	fmt.Println("Perkenalkan saya adalah", nama+", salah satu mahasiswa Prodi S1-TI dari Kelas", kelas, "dengan NIM", nim+".")
}
