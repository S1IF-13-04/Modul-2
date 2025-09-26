package main

import "fmt"

func main() {
	var Nama string
	var Nim int
	var Kelas string

	fmt.Scanln(&Nama, &Nim, &Kelas)
	
	fmt.Println("Perkenalkan saya adalah", Nama,
		"salah satu mahasiswa Prodi S1-IF dari kelas", Kelas,
		"dengan NIM", Nim)
}