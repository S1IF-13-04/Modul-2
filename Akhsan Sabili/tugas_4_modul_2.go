package main

import "fmt"

func main(){
	var F, C float64
	fmt.Print("Masukkan suhu dalam F :")
	fmt.Scanln(&F)
	C = (F-32)*5/9
	fmt.Println(C)
}