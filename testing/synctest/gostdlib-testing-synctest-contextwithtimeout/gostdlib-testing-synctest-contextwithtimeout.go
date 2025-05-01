package main

import (
	"context"
	"fmt"

	//"testing/synctest"
	"time"
)

func withTimeout() {
	// Create a context.Context which is canceled after a timeout.
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Wait just less than the timeout.
	time.Sleep(timeout - time.Nanosecond)
	//synctest.Wait()
	fmt.Printf("before timeout: ctx.Err() = %v\n", ctx.Err())

	// Wait the rest of the way until the timeout.
	time.Sleep(time.Nanosecond)
	//synctest.Wait()
	fmt.Printf("after timeout:  ctx.Err() = %v\n", ctx.Err())

}

func main() {
	//synctest.Run(withTimeout)
	withTimeout()
}
