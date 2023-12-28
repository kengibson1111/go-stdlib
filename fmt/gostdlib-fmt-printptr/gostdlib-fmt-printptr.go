package main

import "fmt"

func main() {
    var sample = true

    fmt.Println("Println:")
    fmt.Println(sample)

    fmt.Println("Printf with %p:")
    fmt.Printf("%p\n", &sample)
}
