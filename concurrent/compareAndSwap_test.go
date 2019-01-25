package main

import (
	"sync"
	"testing"
)

func TestCompareAndSwap(t *testing.T) {
	size := 10
	ts := newTicketStore(size)

	wg := sync.WaitGroup{}
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func(v int) {
			defer wg.Done()
			ts.book("s" + string(v))
		}(i)
		go func() {
			ts.getDone()
		}()
	}

	wg.Wait()

	if *ts.ticket != uint64(size) {
		t.Errorf("Expected 10 tickets, got: %v", *ts.ticket)
	}
	if *ts.done != uint64(size) {
		t.Errorf("Expected 10 dones, got: %v", *ts.done)
	}

	for i := 0; i < size; i++ {
		if ts.store[i][0] != 's' {
			t.Errorf("Expected s, got: %v", ts.store[i][0])

		}
	}
}
