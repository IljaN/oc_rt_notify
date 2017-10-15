package main

import (
	"sync"
)

type Subscribers struct {
	sync.Map
}

func (m *Subscribers) Add(sub *Subscriber) *Subscriber {
	_, ok := m.Get(sub.Id)

	if ok {
		panic("Subscriber does already exist")
	}

	stored, _ := m.LoadOrStore(sub.Id, sub)
	return stored.(*Subscriber)
}

func (m *Subscribers) Get(id string) (*Subscriber, bool) {
	v, ok := m.Load(id)

	if !ok {
		return nil, false
	}

	sub, ok := v.(*Subscriber)
	return sub, ok
}
