package main

import "fmt"

func main() {
	var(
		satu, dua, tiga, string
		temp string
	)
	fmt.Println("masukan input string")	
	fmt.Scan(&satu)	
	fmt.Println("masukan input string")
	fmt.Scan(&dua)
	fmt.Print("masukan input string")
	fmt.Scan(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Ouput akhir = " + satu + "" + dua + " " + tiga)
}	

