package main

import (
    "fmt"
    "my-go-app/pkg"
)

func main() {
    fmt.Println("Welcome to My Go App!")
    result := pkg.Add(5, 3)
    fmt.Printf("The result of addition is: %d\n", result)
}