package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
    const nihongo = "日本語"

    fmt.Printf("plain string: ")
    fmt.Printf("%s", nihongo)
    fmt.Printf("\n")

    fmt.Printf("quoted string: ")
    fmt.Printf("%+q", nihongo)
    fmt.Printf("\n")

    fmt.Printf("runes: ")
    for i, w := 0, 0; i < len(nihongo); i += w {
        runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
        w = width
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
    for i, w := 0, 0; i < len(latin); i += w {
        runeValue, width := utf8.DecodeRuneInString(latin[i:])
        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
        w = width
    }
    fmt.Printf("\n")
}
