package main

import (
	"context"
	"flag"
	"fmt"
	"lab/log"
	"lab/service"
	stdlog "log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	dest := flag.String("dest", "./log.log", "File path where log should be written")
	port := flag.Uint("port", 3000, "TCP port that port should listen on")

	flag.Parse()
	log.Run(*dest)
	service.RegisterHandlers()

	s := http.Server{
		Addr: fmt.Sprintf(":%v", *port),
	}
	go func() {
		fmt.Printf("Logging service started on port %v. Press <enter> to stop.", *port)
		var a string
		fmt.Scanln(&a)
		fmt.Println("Shutting service down.")
		s.Shutdown(context.TODO())
	}()

	stdlog.Println(s.ListenAndServe())

}
