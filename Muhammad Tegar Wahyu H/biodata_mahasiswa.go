package main

import "fmt"

func main() {
	var nama, nim, kelas string

	// Input biodata
	fmt.Print("Masukkan nama mahasiswa: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukkan kelas mahasiswa: ")
	fmt.Scanln(&kelas)

	// Output biodata
	fmt.Printf("\n=== BIODATA MAHASISWA ===\n")
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa\n", nama)
	fmt.Printf("Prodi S1-IF dari kelas %s dengan NIM %s.\n", kelas, nim)
	
	fmt.Printf("\n=== INFORMASI TAMBAHAN ===\n")
	fmt.Printf("Nama Lengkap  : %s\n", nama)
	fmt.Printf("Nomor Induk   : %s\n", nim)
	fmt.Printf("Kelas         : %s\n", kelas)
	fmt.Printf("Program Studi : Informatika (S1)\n")
}