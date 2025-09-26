package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	//namaPanggil := strings.Split(nama, " ")[0]

	fmt.Print("Masukkan NIM: ")
	nim, _ := reader.ReadString('\n')
	nim = strings.TrimSpace(nim)

	fmt.Print("Masukkan kelas: ")
	kelas, _ := reader.ReadString('\n')
	kelas = strings.TrimSpace(kelas)

	Prodi := strings.Split(kelas, "-")[1]
	Prodi = strings.Replace(Prodi, "x", "X", 1)
	Prodi = strings.Replace(Prodi, "X", "", 1)

	fmt.Printf("#BIO-DATA#\nPerkenalkan nama saya adalah %s,\nsalah satu mahasiswa Prodi S1-%s dari kelas %s dengan NIM %s.\n", nama, Prodi, kelas, nim)
}
