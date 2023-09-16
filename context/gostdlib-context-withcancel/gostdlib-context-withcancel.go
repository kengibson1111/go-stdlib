package main

import (
	"context"
	"fmt"
)

// gen() generates integers in a separate goroutine closure
// and sends them via a newly allocated, returned "receive only"
// channel. The caller of gen() uses the channel to "receive only"
// from the channel one integer at a time since the channel is not
// buffered. The caller of gen() need to cancel the context once
// it is done consuming generated integers in order not to leak
// the internal, closure goroutine started by gen().
//
// "receive only" channels cannot be closed explicitly. Go knows when
// the channel is no longer needed.
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1

	go func() {
		for {
			// channels need a valid sender and receiver to not be blocked.
			// ctx.Done() returns a closed channel with a blank struct type.
			// Whenever a blank struct type (struct{}) is sent via channel
			// it consumes zero memory.
			//
			// The syntax "<-channel" can be a way to block code until a struct{}
			// is sent on the channel. When "<-channel" is in a select block,
			// the case blocks, but other cases can activate. This is
			// a notification pattern in Go. In other words, the actual content
			// sent via the channel is not important. The important part is that
			// the channel is activated even if the activation is done with
			// zero bytes.
			//
			// For the second case, if the dst channel has room for a single integer
			// (size of the channel buffer), it will activate. If the channel is full,
			// that means channel receivers haven't received the integer allowing
			// another integer to be sent via the channel. It also means the
			// second case won't activate. The first time in the for loop, an integer
			// will be put on the channel. After that, the second case will not activate
			// until the integer already on the channel has been received.
			//
			// when neither activates, the code blocks.
			select {

			// this will be activated by the "defer cancel()" in main(). The return
			// is how to prevent goroutine leaks.
			case <-ctx.Done():
				return

			// this will be activated in the main() for loop every time an integer is received
			// from the dst channel.
			case dst <- n:
				n++
			}
		}
	}()

	// the dst channel is returned after gen() is called once in the main() for loop. But,
	// gen()'s goroutine is still running.
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// The goroutine closure inside of gen() spins in its for loop until
	// either a ctx.Done() returns a valid channel (and sends a valid struct{})
	// OR a dst channel receiver asks for another integer. When a dst channel
	// receiver asks for another integer, case dst <- n in gen()'s goroutine
	// is activated, another integer is sent, and the goroutine then spins in
	// its for loop.
	dst := gen(ctx)

	for n := range dst {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	// the defer cancel() is really here on the execution path - at the
	// end of the function where the defer is activated. cancel() is going
	// to activate the "case <-ctx.Done()" in gen()'s goroutine.
}
