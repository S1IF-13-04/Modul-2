package main

import "fmt"

func main() {
	var nama, nim, kelas string

	fmt.Println("masukkan nama: ")
	fmt.Scan(&nama)
	fmt.Println("masukkan nim: ")
	fmt.Scan(&nim)
	fmt.Println("masukkan kelas: ")
	fmt.Scan(&kelas)

	fmt.Printf("perkenalkan saya adalah %s salah satu mahasiswa prodi S1-IF dari kelas %s dengan nim %s ", nama, kelas, nim)

}
