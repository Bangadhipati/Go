package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Println("Input a number")
    fmt.Scanln(&n)
    if n >= 18 {
        fmt.Println("You are Eligible for Vote")
    } else {
        fmt.Println("You are not Eligible for Vote")
    }
}
