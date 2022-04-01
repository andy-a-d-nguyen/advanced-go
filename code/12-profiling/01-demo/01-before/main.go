package main

import (
	"context"
	"flag"
	"fmt"
	"lab/log"
	"lab/service"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	dest := flag.String("dest", "./log.log", "File path where log should be written")
	port := flag.Uint("port", 3000, "TCP port that port should listen on")

	flag.Parse()
	log.Destination = *dest
	log.Run()
	service.RegisterHandlers()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
		if err != nil {
			fmt.Println(err)
			cancel()
		}
	}()
	go func() {
		fmt.Printf("Logging service started on port %v. Press any key to stop.", *port)
		var a string
		fmt.Scanln(&a)
		fmt.Println("Shutting service down.")
		cancel()
	}()

	<-ctx.Done()
}
