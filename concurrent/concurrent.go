package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Player ping or pong
type Player string

// Ball tokick from ping an pong and count the hits
type Ball struct {
	hits          int
	currentPlayer Player
}

func (ball Ball) String() string {
	return fmt.Sprintf("[hit: %d from: %s]", ball.hits, ball.currentPlayer)
}

// Dir list file and directories recursive from start diectory
func play(player Player, ballChan chan *Ball) {
	for {
		// receive value from chain
		// other go-routine is blocked and can not receive new data
		ball := <-ballChan
		ball.hits++
		ball.currentPlayer = player
		dur := time.Millisecond * time.Duration(rand.Intn(1000))
		time.Sleep(dur)
		fmt.Printf("Play: %v -- [%v] \n", ball.String(), dur)
		// ready and send result to chain
		ballChan <- ball
	}
}

func run() {
	var (
		ping Player = "ping"
		pong Player = "pong"
	)

	ballChan := make(chan *Ball)

	// start 2 player and wait to send data on channel
	go play(ping, ballChan)
	go play(pong, ballChan)

	// init new Ball and send to Channel
	// !!! block channel !!!
	// receiver must start first, otherwise nobody can't unlock
	ballChan <- new(Ball)
	time.Sleep(time.Second * 3)
	// release channel, the main functions is completed
	<-ballChan
}

func main() {
	run()
}

// Example with WaitGroup and Buffer Channel
func exampleWithBufferChannelAndWaitGroup() {
	// Buffer = 2 (without Buffer, the channel blocked)
	sem := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	sem <- 1
	go func() {
		defer wg.Done()
		fmt.Println("Ja")
		<-sem
	}()

	wg.Wait()

	// wait until press Enter
	// var in byte = '\n'
	// reader := bufio.NewReader(os.Stdin)
	// reader.ReadString(in)

}

// To Examples to illustrate the Problem vom blocking Channel:
// ------------------------------------------------------------

// first: start go-function (in background) with receive value from channel
// secend: send value to channel
func simpleExampleWork() {
	c := make(chan int)
	fmt.Println("start")

	go func() {
		fmt.Println("receive: ", <-c)
	}()
	c <- 1
}

// deadlock - send before can't receive value from channel
func simpleExampleDoNotWork() {
	c := make(chan int)
	fmt.Println("start")
	// blocked !!!
	c <- 1
	// this go function was NEVER reached
	go func() {
		fmt.Println("receive: ", <-c)
	}()
}
