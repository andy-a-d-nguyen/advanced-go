package main

import (
	"fmt"
	"sync"
)

func main() {
	s := make([]int, 0)
	wg := sync.WaitGroup{}
	const iterations = 10000
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			s = append(s, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(len(s))
}
