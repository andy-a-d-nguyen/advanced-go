package alog

import (
	"log"
	"os"
)

var (
	destination string = "./application.log"
	lw          logWriter
)

func Run(dest string) {
	dest = destination
	log.SetOutput(&lw)
}

type logWriter struct{}

func (logWriter) Write(msg []byte) (int, error) {
	f, err := os.OpenFile(destination, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0200)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(msg)
}
