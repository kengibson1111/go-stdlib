package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f\n", math.Sin(math.Pi))
	fmt.Printf("%.2f\n", math.Sinh(0))

	sin, cos := math.Sincos(0)
	fmt.Printf("%.2f, %.2f\n", sin, cos)
}
