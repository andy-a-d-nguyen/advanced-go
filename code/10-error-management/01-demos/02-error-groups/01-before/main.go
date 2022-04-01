package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	resultCh := make(chan string)
	errCh := make(chan error)
	wg := new(sync.WaitGroup)

	const start = `c:\go`

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wg.Add(1)
	go listFiles(ctx, start, resultCh, errCh, wg)

	go func() {
		err := <-errCh
		cancel()
		log.Fatal("Canceled due to error: ", err)
	}()

	go func() {
	loop:
		for {
			select {
			case path := <-resultCh:
				fmt.Println(path)
			case <-ctx.Done():
				log.Fatal(ctx.Err())
				break loop
			}
		}
	}()
	wg.Wait()

}

func listFiles(ctx context.Context, path string, resultCh chan<- string, errCh chan<- error, wg *sync.WaitGroup) {
	f, err := os.Open(path)
	if err != nil {
		errCh <- err
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		errCh <- err
		return
	}
	if fi.IsDir() {
		select {
		case <-ctx.Done():
			return
		default:
		}
		children, err := f.Readdir(0)
		if err != nil {
			errCh <- err
			return
		}
		wg.Add(len(children))
		for _, c := range children {
			go listFiles(ctx, filepath.Join(path, c.Name()), resultCh, errCh, wg)
		}
	} else {
		resultCh <- path
	}
	wg.Done()

}
