package main

import (
	"runtime"
	"sync/atomic"
)

// a youtube-video describe this implementation:
// https://www.youtube.com/watch?v=YEKjSzIwAdA&t=305s

type ticketStore struct {
	ticket *uint64
	done   *uint64
	store  []string
}

func newTicketStore(size int) *ticketStore {
	return &ticketStore{
		ticket: new(uint64),
		done:   new(uint64),
		store:  make([]string, size+1),
	}
}

func (ts *ticketStore) book(s string) {
	// increment global variable ticket
	// and substract -1 localy
	// draw a ticket (increment ticket, pos in store = ticket - 1)
	t := atomic.AddUint64(ts.ticket, 1) - 1
	ts.store[t] = s
	// increment done thread safe
	for !atomic.CompareAndSwapUint64(ts.done, t, t+1) {
		runtime.Gosched()
	}
}

func (ts *ticketStore) getDone() []string {
	return ts.store[:atomic.LoadUint64(ts.done)+1]
}
