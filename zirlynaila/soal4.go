// package main
// import "fmt"

// func main() {
// 	var fahreinheit, celcius int
// 	fmt.Print("Masukkan suhu dalam Fahreinheit= ")
// 	fmt.Scanln(&fahreinheit)
// 	celcius = (fahreinheit - 32) * 5 / 9
// 	fmt.Println("Suhu dalam Celcius= ", celcius)
// }

package main
import "fmt"

func main() {
    var nama string
    fmt.Print("Masukkan nama anda= ")
    fmt.Scanln(&nama)
    fmt.Println("nama saya adalah ", nama)
}
