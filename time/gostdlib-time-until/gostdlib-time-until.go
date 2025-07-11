package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	futureTime := time.Now().Add(5 * time.Second)
	durationUntil := time.Until(futureTime)

	fmt.Printf("Duration until future time: %.0f seconds\n", math.Ceil(durationUntil.Seconds()))
}
