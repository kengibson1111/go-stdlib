package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	positiveDuration := 5 * time.Second
	negativeDuration := -3 * time.Second
	minInt64CaseDuration := time.Duration(math.MinInt64)

	absPositive := positiveDuration.Abs()
	absNegative := negativeDuration.Abs()
	absSpecial := minInt64CaseDuration.Abs() == time.Duration(math.MaxInt64)

	fmt.Printf("Absolute value of positive duration: %v\n", absPositive)
	fmt.Printf("Absolute value of negative duration: %v\n", absNegative)
	fmt.Printf("Absolute value of MinInt64 equal to MaxInt64: %t\n", absSpecial)
}
