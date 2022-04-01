// Package channel provides an enhanced version of the built-in
// buffered channel. It adds the ability to monitor the
// number of messages in the buffer as well as the open/closed
// status of the channel.
package schannel

import (
	"errors"
	"sync"
)

type smartChannel struct {
	closed       bool
	ch           chan interface{}
	bufferDepth  int
	m            *sync.Mutex
	messageCount int
}

type SmartChannel interface {
	Close() bool
	Send(msg interface{}) error
	Receive() (interface{}, bool)
	IsClosed() bool
	MessageCount() int
}

func New(bufferSize int) SmartChannel {
	return &smartChannel{
		ch: make(chan interface{}, bufferSize),
		m:  new(sync.Mutex),
	}
}

func (sc *smartChannel) Close() bool {
	close(sc.ch)
	sc.closed = true
	return sc.closed
}

func (sc *smartChannel) Send(msg interface{}) error {
	if sc.closed {
		return errors.New("Send on closed channel")
	}
	sc.m.Lock()
	defer func() {
		sc.m.Unlock()
	}()
	defer func() {
		// assume all failures are due to closed channel
		if recover() != nil {
			sc.closed = true
		} else {

			sc.messageCount++
		}
	}()
	sc.ch <- msg
	return nil
}

func (sc *smartChannel) Receive() (interface{}, bool) {
	sc.m.Lock()
	defer func() {
		sc.m.Unlock()
	}()
	msg, ok := <-sc.ch
	if ok {
		sc.messageCount--
	}
	return msg, ok
}

func (sc smartChannel) IsClosed() bool {
	return sc.closed
}

func (sc smartChannel) MessageCount() int {
	sc.m.Lock()
	defer sc.m.Unlock()
	return sc.messageCount
}
