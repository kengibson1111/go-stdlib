package main

import (
	"fmt"
	"math"
)

func main() {
	p := math.Round(10.5)
	fmt.Printf("%.1f\n", p)

	n := math.Round(-10.5)
	fmt.Printf("%.1f\n", n)

	u := math.RoundToEven(11.5)
	fmt.Printf("%.1f\n", u)

	d := math.RoundToEven(12.5)
	fmt.Printf("%.1f\n", d)
}
