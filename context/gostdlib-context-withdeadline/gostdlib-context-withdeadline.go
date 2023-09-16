package main

import (
	"context"
	"fmt"
	"time"
)

// general-purpose function that will simulate a processing
// delay and return an integer on the reply channel. The channel
// is a "receive only" channel. The arrow is on the left side of
// the chan keyword. The channel type can be anything, but for
// this example the type is a simple integer.
func processingFunc(delay int, replyChannel chan<- int) {
	fmt.Println("entering processingFunc...")

	// delay to simulate processing
	time.Sleep(time.Second * time.Duration(delay))
	replyChannel <- 0

	fmt.Println("leaving processingFunc...")

	// clean up any resources using defer statements in processing code.
}

func deadlineAlwaysReached() {
	// a blank struct type channel is an implementation of a notification
	// pattern in Go. In other words, the actual content sent via the
	// channel is not important. The important part is that the channel is
	// activated even if the activation is done with zero bytes.
	neverUsed := make(chan struct{})
	defer close(neverUsed)

	// the deadline can be any future date. for this example, it's
	// simple to add 4 seconds to current time.
	d := time.Now().Add(time.Second * 4)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	// code blocks at the select until one of the cases activates.
	select {
	// this will never activate because the neverUsed channel is not closed
	// until after the select (deferred).
	case <-neverUsed:
		fmt.Println("shouldn't get here")

	// this will activate
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	// neverUsed closed and context cancelled based on defers.
}

func processingCompletesBeforeDeadline() {
	// make the channel for the processing reply and defer the close.
	replyChannel := make(chan int)
	defer close(replyChannel)

	// the deadline can be any future date. for this example, it's
	// simple to add 4 seconds to current time.
	d := time.Now().Add(time.Second * 4)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	// processing delay will result in not reaching the deadline.
	// processingFunc is activated as a separate goroutine (thread).
	go processingFunc(3, replyChannel)

	// code blocks at the select until one of the cases activates.
	select {
	// this should activate because the processing delay is
	// not big enough to result in a deadline condition.
	case replyInt := <-replyChannel:
		fmt.Println("processing reply =", replyInt)

	// this should not activate
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	// replyChannel closed and context cancelled based on defers.
}

func processingCompletesAfterDeadline() {
	// make the channel for the processing reply and defer the close.
	replyChannel := make(chan int)
	defer close(replyChannel)

	// the deadline can be any future date. for this example, it's
	// simple to add 4 seconds to current time.
	d := time.Now().Add(time.Second * 4)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	// processing delay will result in reaching the deadline.
	// processingFunc is activated as a separate goroutine (thread).
	go processingFunc(5, replyChannel)

	// code blocks at the select until one of the cases activates.
	select {
	// this should not activate because the processing delay is
	// big enough to result in a deadline condition.
	case replyInt := <-replyChannel:
		fmt.Println("processing reply =", replyInt)

	// this should activate
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	// replyChannel closed and context cancelled based on defers.
}

func main() {
	// this is the simple case with a channel that is never used and never
	// closed until defer logic.
	fmt.Println("deadlineAlwaysReached...")
	deadlineAlwaysReached()

	// this is a simple case when processing completes before the deadline
	// is reached.
	fmt.Println("\nprocessingCompletesBeforeDeadline...")
	processingCompletesBeforeDeadline()

	// this is a simple case when processing completes after the deadline
	// is reached. Notice main() completes before processingFunc() has
	// a chance to do its last print. In the case of server processing,
	// it could print somewhere because a server isn't meant to shut down
	// like a short-lived main() example.
	fmt.Println("\nprocessingCompletesAfterDeadline...")
	processingCompletesAfterDeadline()
}
