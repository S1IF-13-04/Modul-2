package main

import "fmt"
		
func main (){
	var p float64
	fmt.Printf("Masukan Jari-jari : ")
	fmt.Scan(&p)
	luas := 3.14 * p * p
	fmt.Printf("hasil : %.1f\n",luas)
}