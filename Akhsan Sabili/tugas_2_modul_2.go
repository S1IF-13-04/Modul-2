package main

import "fmt"

func main(){
	var nama, kelas, nim string
	
	fmt.Print("nama :")
	fmt.Scanln(&nama)
	fmt.Print("kelas :")
	fmt.Scanln(&kelas)
	fmt.Print("NIM :")
	fmt.Scanln(&nim)
	fmt.Printf("Perkenalkan saya adalah %s salah satu mahasiswa prodi S1-IF dari kelas %s dengan NIM %s \n", nama, kelas, nim )
	
}