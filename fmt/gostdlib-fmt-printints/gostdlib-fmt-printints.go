package main

import "fmt"

func main() {
    const sample = 1234

    fmt.Println("Println:")
    fmt.Println(sample)

    fmt.Println("Printf with %b:")
    fmt.Printf("%b\n", sample)

    fmt.Println("Printf with %c:")
    fmt.Printf("%c\n", sample)

    fmt.Println("Printf with %d:")
    fmt.Printf("%d\n", sample)

    fmt.Println("Printf with %o:")
    fmt.Printf("%o\n", sample)

    fmt.Println("Printf with %q:")
    fmt.Printf("%q\n", sample)

    fmt.Println("Printf with %x:")
    fmt.Printf("%x\n", sample)

    fmt.Println("Printf with %X:")
    fmt.Printf("%X\n", sample)

    fmt.Println("Printf with %U:")
    fmt.Printf("%U\n", sample)
}
