package main

import (
	"github.com/nats-io/go-nats-streaming"
	"sync"
)

type Session struct {
	Id         string
	Messages   chan *stan.Msg
	Subscriber *Subscriber
	closed     bool
	mutex      sync.Mutex
}

func (ses *Session) SafeClose() {
	ses.mutex.Lock()
	if !ses.closed {
		close(ses.Messages)
		ses.closed = true
	}
	ses.mutex.Unlock()
}

func (ses *Session) IsClosed() bool {
	ses.mutex.Lock()
	defer ses.mutex.Unlock()
	return ses.closed
}
