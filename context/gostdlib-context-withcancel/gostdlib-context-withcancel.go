package main

import (
	"context"
	"fmt"
)

func main() {
	// gen() generates integers in a separate goroutine closure
	// and sends them to a newly allocated, returned channel.
	// The caller of gen() need to cancel the context once
	// it is done consuming generated integers in order not to leak
	// the internal, closure goroutine started by gen().
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1

		go func() {
			for {
				// channels need a valid sender and receiver to activate.
				// if the first case doesn't have a valid channel from ctx.Done()
				// yet, it won't activate because there is no channel and message sent
				// on the channel. ctx.Done() does both of those things at the same time.
				// If the second case doesn't have a dst channel receiver yet, it won't
				// activate. The channel and the message sent are valid, but no channel
				// receiver.
				//
				// when neither activate, the code waits at the select.
				select {

				// this will be activated by the "defer cancel()" in main(). The return
				// is how to prevent goroutine leaks.
				case <-ctx.Done():
					return

				// this will be activated in the main() for loop call to gen(). The for loop
				// iterates every time a new value appears on the channel.
				case dst <- n:
					n++
				}
			}
		}()

		// the dst channel is returned after gen() is called once in the main() for loop. But,
		// gen()'s goroutine is still running.
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// even though the for loop is executing 5 times, gen() is only
	// called once returning an active channel that is used 5 times.
	// The goroutine closure inside of gen() waits at the select until
	// either a ctx.Done() returns a valid channel (and sends a valid struct{})
	// OR a dst channel receiver asks for more. When a dst channel receiver
	// asks for more, case dst <- n in gen()'s goroutine is activated, another
	// integer is sent, and execution then waits at the select again.
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	// the defer cancel() is really here on the execution path - at the
	// end of the function where the defer is activated. cancel() is going
	// to activate the "case <-ctx.Done()" in gen()'s goroutine.
}
