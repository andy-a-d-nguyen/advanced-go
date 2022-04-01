package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	/* ---------------GOROUTINE-----------------*/

	// WaitGroup is a thread-safe counter
	// var wg sync.WaitGroup

	// wg.Add(1)
	// creating a goroutine
	// instructions get passed to the Go scheduler, which asks for resource from OS threads + create a virtual thread
	// go func() {
	// 	fmt.Println("Hello World")
	// 	wg.Done()
	// }()

	// this sees counter from .Add() and wait for .Done() to be called
	// .Done() decrements the thread counter
	// wg.Wait()

	// Concurrency: Things CAN happen at the same time
	// Parallelism: Things ARE happening at the same time

	/* ---------------CHANNEL-----------------*/

	// Channels: Pipelines that communicate between goroutines

	// var wg sync.WaitGroup

	// to make a channel
	// type of channel: type of info to bet passed between goroutines
	// ch := make(chan string)

	// wg.Add(2)

	// to prevent race condition, pass closure variables to goroutines
	// and make channel parameter unidirectional with "<-"
	// go func(ch chan<- string, wg *sync.WaitGroup) {
	// 	fmt.Println("one")
	// let message flow INTO channel
	// 	ch <- "The message"
	// 	wg.Done()
	// }(ch, &wg)

	// go func(ch <-chan string, wg *sync.WaitGroup) {
	// 	fmt.Println("two")
	// let message flow OUT OF channel
	// 	fmt.Println(<-ch)
	// 	wg.Done()
	// }(ch, &wg)

	// wg.Wait()

	// goroutines may be created out of order, but the go scheduler will know which one is trying to receive data from another goroutine
	// the channel is the mechanism in which data between goroutines are synchronized

	// go passes data by value unless specified to pass pointer

	/* --------------------------------*/

	// var wg sync.WaitGroup
	// ch := make(chan string)

	// wg.Add(10)

	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		ch <- fmt.Sprintf("Current counter val: %v\n", i)
	// 		wg.Done()
	// 	}(i)
	// }

	// In order to deal with multiple producers in goroutines, create a goroutine that acts as a supervisor to watch for count of WaitGroup
	// go func() {
	// wg.Wait()
	// close() closes the sending side of a channel
	// close(ch)
	// }()

	// a channel is always open and therefore, it must be closed when we are done to avoid deadlocks
	// channels have to be read from before Wait() can be called; if not, a deadlock will happen
	// The for … range construct iterates until the channel is closed. It doesn't process that closure until the last message has been processed, so it handles all 10 messages, detects the closed channel, and terminates the loop.
	// if using a regular for loop, now the receiving side knows how many messages to read from
	// for msg := range ch {
	// fmt.Println(msg)
	// }

	// wg.Wait()

	/* ---------------CONTEXT-----------------*/

	// A context can tell other goroutines to terminate and is shared between goroutines
	// context can be passed along between families of goroutines (between parents and children)
	// when a context is cancelled, all the context that's shared is cancelled
	// when a parent context is cancelled, all children contexts are cancelled

	var wg sync.WaitGroup
	ch := make(chan string)

	// Background() has no constraints
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer cancel()

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(ctx context.Context, i int) {
			// if context is cancelled, don't do any work
			if ctx.Err() != nil {
				wg.Done()
			}
			ch <- fmt.Sprintf("Current counter val: %v\n", i)
			wg.Done()
		}(ctx, i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for {
		// select reads from multiple channels
		select {
		// second var indicates whether channel is open
		case msg, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(msg)
		// Done() returns a closed channel when context is closed
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}
