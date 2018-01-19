package main

import "testing"

// HINT: you can run the tests with: go test -race ./...

func TestPlayer(t *testing.T) {
	var (
		ping Player = "ping"
	)

	ballChan := make(chan *Ball)
	go play(ping, ballChan)
	ballChan <- new(Ball)
	ball := <-ballChan

	t.Logf("Player: %v Ball-CurrentPlayer: %v", ping, ball.currentPlayer)
	if ball.currentPlayer != ping {
		t.Errorf("Expected Player %v, but was: %v", ping, ball.currentPlayer)
	}
}
