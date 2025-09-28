package main
import "fmt"

func main() {
	var c1, c2, c3, c4, c5 byte
	var b1, b2, b3 byte

	fmt.Scan(&c1, &c2, &c3, &c4, &c5)

	var str string
	fmt.Scan(&str)

	b1 = str[0]
	b2 = str[1]
	b3 = str[2]

	fmt.Printf("%c%c%c%c%c\n", c1, c2, c3, c4, c5) 
	fmt.Printf("%c%c%c", b1+1, b2+1, b3+1)
}
