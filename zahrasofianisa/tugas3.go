package main
import "fmt"

func main() {
	const phi float64 = 3.14159
	var r float64
	fmt.Scan(&r)
	hasil := phi * (r * r)
	fmt.Print(hasil)
}

