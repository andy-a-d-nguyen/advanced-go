package main

import (
	"demo/schannel"
	"fmt"
	"log"
	"sync"
)

func main() {
	ch := schannel.New(3)

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			err := ch.Send(i)
			fmt.Println("Message count:", ch.MessageCount())
			if err != nil {
				log.Print(err)
			}
		}(i)
		go func() {
			msg, _ := ch.Receive()
			fmt.Println(msg)
			wg.Done()
		}()
	}
	wg.Wait()

	ch.Close()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			err := ch.Send(i)
			if err != nil {
				log.Print(err)
			}
		}(i)
		go func() {
			msg, _ := ch.Receive()
			fmt.Println(msg)
			wg.Done()
		}()
	}
	wg.Wait()
}
