package main

import "fmt"

func main() {
    const sample = true

    fmt.Println("Println:")
    fmt.Println(sample)

    fmt.Println("Printf with %t:")
    fmt.Printf("%t\n", sample)
}
