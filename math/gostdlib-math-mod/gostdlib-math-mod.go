package main

import (
	"fmt"
	"math"
)

func main() {
	c := math.Mod(7, 4)
	fmt.Printf("%.1f\n", c)

	int, frac := math.Modf(3.14)
	fmt.Printf("%.2f, %.2f\n", int, frac)

	int, frac = math.Modf(-2.71)
	fmt.Printf("%.2f, %.2f\n", int, frac)
}
