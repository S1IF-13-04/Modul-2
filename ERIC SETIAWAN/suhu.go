package main

import (
	"fmt"
)

func main() {
	var f float64

	fmt.Scanln(&f)

	c := (f - 32) * 5 / 9

	fmt.Printf("%.0f\n", c)
}
