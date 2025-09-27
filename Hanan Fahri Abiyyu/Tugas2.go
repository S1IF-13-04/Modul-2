package main

import "fmt"

func main () {

	var nama, nim, kelas string

	fmt.Print("Masukkan nama : ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan NIM : ")
	fmt.Scanln(&nim)
	fmt.Print("Masukkan kelas : ")
	fmt.Scanln(&kelas)

	fmt.Printf("Halo, Perkenalkan nama saya %s salah satu mahasiswa prodi Teknik Informatika, kelas %s, NIM %s. Hobi saya adalah bermain Bola Basket", nama, kelas, nim)
}
