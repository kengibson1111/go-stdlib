package main

import (
	"fmt"
	"time"
)

func expensiveCall() {}

func main() {
	start := time.Now()
	expensiveCall()
	elapsed := time.Since(start)

	fmt.Printf("The call took %v to run.\n", elapsed)
}
