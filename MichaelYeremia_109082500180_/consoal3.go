package main

import "fmt"

func main() {
    var nums [11]int
    for i := 0; i < 11; i++ {
        fmt.Scan(&nums[i])
    }

    for i := 0; i < 11; i++ {
        fmt.Printf("%c", nums[i])
    }
    fmt.Println() // pindah baris

    var chars [8]rune
    for i := 0; i < 8; i++ {
        fmt.Scanf("%c", &chars[i])
    }

    for i := 0; i < 8; i++ {
        nextChar := chars[i] + 1
        fmt.Printf("%c", nextChar)
    }
    fmt.Println()
}
