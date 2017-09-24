package service

import (
	"time"
	"math/rand"
)

type RandomNumbers struct {
	Numbers chan int
}

func NewBeacon() *RandomNumbers {
	rnd := &RandomNumbers{
		Numbers: make(chan int, 20),
	}

	go rnd.Start()
	return rnd
}

func (rnd *RandomNumbers) Start() {
	go func() {
		for {
			rnd.Numbers <- rand.Int()
			time.Sleep(1 * time.Second)
		}

	}()

}

