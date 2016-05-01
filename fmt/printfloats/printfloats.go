package main

import "fmt"

func main() {
    const sample = 12345.6789

    fmt.Println("Println:")
    fmt.Println(sample)

    fmt.Println("Printf with %b:")
    fmt.Printf("%b\n", sample)

    fmt.Println("Printf with %e:")
    fmt.Printf("%e\n", sample)

    fmt.Println("Printf with %E:")
    fmt.Printf("%E\n", sample)

    fmt.Println("Printf with %f:")
    fmt.Printf("%f\n", sample)

    fmt.Println("Printf with %F:")
    fmt.Printf("%F\n", sample)

    fmt.Println("Printf with %g:")
    fmt.Printf("%g\n", sample)

    fmt.Println("Printf with %G:")
    fmt.Printf("%G\n", sample)
}
