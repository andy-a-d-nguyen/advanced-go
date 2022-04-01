// Package main docs.
package demo

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	resultCh := make(chan string)
	const start = `c:\go`

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(listFiles(ctx, start, resultCh, g))

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
	err := g.Wait()
	if err != nil {
		log.Fatal(err)
	}

}

func listFiles(ctx context.Context, path string, resultCh chan<- string, g *errgroup.Group) func() error {
	return func() error {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		fi, err := f.Stat()
		if err != nil {
			return err
		}
		if fi.IsDir() {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			children, err := f.Readdir(0)
			if err != nil {
				return err
			}
			for _, c := range children {
				g.Go(listFiles(ctx, filepath.Join(path, c.Name()), resultCh, g))
			}
		} else {
			resultCh <- path
		}
		return nil
	}
}
