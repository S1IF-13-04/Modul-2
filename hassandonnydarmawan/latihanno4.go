package main
import "fmt"

func main () {
	var f,c int64
	fmt.Print("masukan suhu dalam fahrenheit: ")
	fmt.Scan(&f)
	c = (f - 32) * 5/9
	fmt.Println("maka derajat celciusnya adalah", c)
}