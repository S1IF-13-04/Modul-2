package main

import(
	"fmt"
	"math"
)

func main(){
	var r float64

	fmt.Print("Masukan jari jari lingkaran=")
	fmt.Scan(&r)

	luas := math.Pi * r * r
	fmt.Printf("%.1f\n " , luas)
}