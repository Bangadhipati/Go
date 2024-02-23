package main

import "fmt"
func main() {
    var n, c int
    fmt.Println("Input a Number:")
    fmt.Scanln(&n)
    for i := 1; i <= n; i++ {
        if n%i == 0 {
            c++
        }
    }
    if c == 2 {
        fmt.Println("Prime")
    } else {
        fmt.Println("Not Prime")
    }
}