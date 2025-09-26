package main
import "fmt"

func main(){
	var nama, nim, kelas string

	fmt.Print("Masukan nama:")
	fmt.Scanln(&nama)
	fmt.Print("Masukan nim:")
	fmt.Scanln(&nim)
	fmt.Print("Masukan Kelas:")
	fmt.Scanln(&kelas)

	fmt.Println("Perkenalkan saya "+ nama + " Salah satu mahasiswa Prodi S1-IF dari kelas "+ kelas + " dengan nim "+ nim)
	
}