package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// HINT: you can run the tests with: go test -race ./...

func TestPlayer(t *testing.T) {
	var (
		ping Player = "ping"
	)

	ballChan := make(chan Ball)
	go play(ping, ballChan)
	ballChan <- Ball{hits: 0, currentPlayer: ""}
	ball := <-ballChan

	Convey(`Compare Player: "ping" with ball.currentPlayer`, t, func() {
		So(ping, ShouldEqual, ball.currentPlayer)
	})
}
