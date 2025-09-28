package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Scanln(&r)

	luas := math.Pi * r * r

	fmt.Printf("%.1f\n", luas)
}
