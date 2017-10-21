package main

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/nats-io/go-nats-streaming"
	"sync"
)

type Subscriber struct {
	Id            string
	BusConn       stan.Conn
	Subscription  stan.Subscription
	Sessions      *sync.Map
	SessionsCount int
	Broadcast     chan *stan.Msg
}

func NewSubscriber(id string) *Subscriber {
	sub := new(Subscriber)
	sub.Id = id
	sub.Sessions = new(sync.Map)
	con, err := stan.Connect("test-cluster", id)

	if err != nil {
		panic(err)
	}
	sub.BusConn = con
	sub.Broadcast = make(chan *stan.Msg, 100)

	return sub
}

func (sub *Subscriber) HandleBroadcast() {
	go func() {
		for msg := range sub.Broadcast {
			sub.Sessions.Range(func(id, v interface{}) bool {
				go func() {
					ses := v.(*Session)
					if !ses.IsClosed() {
						ses.Messages <- msg
					}
				}()

				return true
			})
		}
	}()
}

func (sub *Subscriber) RegisterSession() *Session {
	buf := make([]byte, 16)
	rand.Read(buf)
	sessid := hex.EncodeToString(buf)

	cs := &Session{
		Id:         sessid,
		Messages:   make(chan *stan.Msg, 50),
		Subscriber: sub,
	}

	v, _ := sub.Sessions.LoadOrStore(sessid, cs)
	ses := v.(*Session)

	if sub.SessionsCount == 0 {
		subscription, err := sub.BusConn.Subscribe(sub.Id, func(msg *stan.Msg) {
			sub.Broadcast <- msg
		}, stan.DurableName("events:"+sub.Id), stan.StartWithLastReceived())

		if err != nil {
			panic(err)
		}

		sub.Subscription = subscription
	}

	sub.SessionsCount++

	return ses

}
