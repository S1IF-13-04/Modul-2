package main

import (
	"fmt"
	"strings"
)

func tentukanFakultas(prodi string) string {
	prodiBersih := strings.ToLower(prodi)

	switch prodiBersih {
	case "informatika", "sistem informasi", "teknik komputer":
		return "Fakultas Informatika"
	case "akuntansi", "manajemen", "ekonomi pembangunan":
		return "Fakultas Ekonomi dan Bisnis"
	case "teknik elektro", "teknik fisika", "teknik telekomunikasi":
		return "Fakultas Teknik Elektro"
	case "ilmu komunikasi", "desain komunikasi visual", "hubungan internasional":
		return "Fakultas Komunikasi dan Industri Kreatif"
	default:
		return "Fakultas (Tidak Dikenal)"
	}
}

func main() {
	var (
		nama string
		kelas string
		prodi string
		fakultas string 
		nim int
	)


	// 1. Proses Input Data
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan Prodi: ")
	fmt.Scanln(&prodi)

	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim)

	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&kelas)

	fakultas = tentukanFakultas(prodi)

	fmt.Println("Perkenalkan nama saya adalah " + nama + ", Saya dari " + fakultas + " Program Study " + prodi + " Dari kelas " + kelas + " dengan NIM ",nim)
	fmt.Println("--------------------------------")
}