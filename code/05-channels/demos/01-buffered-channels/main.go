package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(ch chan<- int) {
		start := time.Now()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		dur := time.Now().Sub(start)
		fmt.Printf("Elapsed time: %v", dur)
		close(ch)
	}(ch)
	for msg := range ch {
		fmt.Println(msg)
		time.Sleep(10 * time.Millisecond)
	}
}
