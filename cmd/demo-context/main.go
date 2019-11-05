package main

import (
	"context"
	"log"
	"time"
)

/**
https://golang.org/pkg/context/#pkg-examples

*/

func slowOperation(ctx context.Context) error {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}

func context_timeout() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := slowOperation(ctx); err != nil {
		log.Fatal("cantext error ",ctx.Err())
	}
}
func main() {
	context_timeout()
}
