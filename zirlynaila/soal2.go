package main
import "fmt"

func main() {
	var nama, kelas, nim string
	fmt.Print("Masukkan nama:")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan nim:")
	fmt.Scanln(&nim)
	fmt.Print("Masukkan kelas:")
	fmt.Scanln(&kelas)
    fmt.Println("Perkenalkan saya adalah",nama,"salah satu mahasiswa Prodi S1-IF dari Kelas",kelas,"dengan NIM",nim)
}