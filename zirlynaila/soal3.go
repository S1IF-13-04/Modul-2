package main
import "fmt"

func main() {
	var jariJari float32
	fmt.Print("Masukkan jari-jari lingkaran= ")
	fmt.Scanln(&jariJari)
	luas := 3.14159 * jariJari * jariJari
	fmt.Printf("Luas lingkaran= %.1f\n", luas)
}