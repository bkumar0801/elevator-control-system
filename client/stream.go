package client

import (
	"log"
	"time"
)

/*
Subscription ...
*/
type Subscription interface {
	Updates() <-chan Request
	Close()
}

/*
Subscribe ...
*/
func Subscribe(device Component) Subscription {
	s := &sub{
		device:  device,
		updates: make(chan Request),
	}
	go s.action()
	return s
}

type sub struct {
	device  Component
	updates chan Request
}

func (s *sub) Updates() <-chan Request {
	return s.updates
}

func (s *sub) Close() {
	close(s.updates)
}

func (s *sub) action() {
	select {
	case s.updates <- s.device.RequestPickup():
	case <-time.After(300 * time.Millisecond):
		log.Println("Sending request timed out!")
	}
}

type merge struct {
	subs    []Subscription
	updates chan Request
}

/*
Merge ...
*/
func Merge(subs ...Subscription) Subscription {
	m := &merge{
		subs:    subs,
		updates: make(chan Request),
	}
	for _, sub := range subs {
		go func(s Subscription) {
			var it Request
			select {
			case it = <-s.Updates():
			}
			select {
			case m.updates <- it:
			}
		}(sub)
	}
	return m
}

func (m *merge) Updates() <-chan Request {
	return m.updates
}

func (m *merge) Close() {
	for _, sub := range m.subs {
		sub.Close()
	}
	close(m.updates)
}
