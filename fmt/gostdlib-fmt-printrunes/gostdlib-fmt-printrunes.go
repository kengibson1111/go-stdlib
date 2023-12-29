package main

import "fmt"

func main() {
    const nihongo = "日本語"

    fmt.Printf("plain string: ")
    fmt.Printf("%s", nihongo)
    fmt.Printf("\n")

    fmt.Printf("quoted string: ")
    fmt.Printf("%+q", nihongo)
    fmt.Printf("\n")

    fmt.Printf("runes: ")
    for index, runeValue := range nihongo {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }
    fmt.Printf("\n")

	const latin = "hello"

    fmt.Printf("plain string: ")
    fmt.Printf("%s", latin)
    fmt.Printf("\n")

    fmt.Printf("quoted string: ")
    fmt.Printf("%+q", latin)
    fmt.Printf("\n")

    fmt.Printf("runes: ")
    for index, runeValue := range latin {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }
    fmt.Printf("\n")
}
