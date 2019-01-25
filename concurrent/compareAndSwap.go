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

func (ts *ticketStore) book(s string) {
	t := atomic.AddUint64(ts.ticket, 1) - 1 // draw a ticket (increment ticket, pos in store = ticket - 1)
	ts.store[t] = s                         // store the data (s)
	for !atomic.CompareAndSwapUint64(ts.done, t, t+1) {
		runtime.Gosched()
	}
}

func (ts *ticketStore) getDone() []string {
	return ts.store[:atomic.LoadUint64(ts.done)+1]
}
