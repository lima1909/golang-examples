package main

import (
	"bufio"
	"context"
	"os"
	"time"

	"fmt"
)

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// cancel this func AND sleepAndTalk with Enter-Key
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	sleepAndTalk(ctx, time.Second*5, "Hello")
}
