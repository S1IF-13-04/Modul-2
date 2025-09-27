package main

import (
	"fmt"
	"math"
)

func main(){
	var r, L float64
	fmt.Print("jari jari :")
	fmt.Scan(&r)
	L = math.Pi * r * r
	fmt.Printf("%.1f\n", L)
}