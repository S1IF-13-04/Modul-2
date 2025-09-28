package main

import "fmt"

func main() {
	var Nama, Kelas string
	var Nim int

	fmt.Scanln(&Nama, &Nim, &Kelas)
	fmt.Println("Perkenalkan saya adalah", Nama,"salah satu mahasiswa Prodi dari kelas", Kelas,"dengan NIM", Nim)
}