package main
import "fmt"

func main() {
	var c1, c2, c3, c4, c5 string
	var str string

	fmt.Scan(&c1, &c2, &c3, &c4, &c5)
	fmt.Scan(&str)

	fmt.Printf("%s%s%s%s%s\n", c1, c2, c3, c4, c5)
	fmt.Printf("%c%c%c\n", str[0]+1, str[1]+1, str[2]+1)
}
