package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type RingBuffer struct {
	inputChannel  <-chan int
	outputChannel chan int
}

func NewRingBuffer(inputChannel <-chan int, outputChannel chan int) *RingBuffer {
	return &RingBuffer{inputChannel, outputChannel}
}

func (r *RingBuffer) Run() {
	for v := range r.inputChannel {
		select {
		case r.outputChannel <- v:
		default:
			<-r.outputChannel
			r.outputChannel <- v
		}
	}
	close(r.outputChannel)
}
func main() {
	in := make(chan int)
	out := make(chan int, 5)
	rb := NewRingBuffer(in, out)
	go rb.Run()

	for i := 0; i < 10; i++ {
		in <- i
	}

	close(in)

	for res := range out {
		fmt.Println(res)
	}
}

////////////////////////////////////////////////////////
func exampleMultiReader() {
	header := strings.NewReader("<msg>")
	msg := strings.NewReader("Hello Mario")
	footer := strings.NewReader("</msg>")
	r := io.MultiReader(header, msg, footer)

	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}

}

func examplePipe() {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		_, err := fmt.Fprintln(pw, "Hello Mario")
		if err != nil {
			panic(err)
		}

	}()

	_, err := io.Copy(os.Stdout, pr)
	if err != nil {
		panic(err)
	}

}
