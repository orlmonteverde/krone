package main

import (
	"context"
	"fmt"
	"time"

	"github.com/orlmonteverde/krone"
)

func main() {
	// context with 10 seconds timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// new krone instance with context.
	k := krone.NewWithContext(ctx, 2*time.Second)

	var count int

	// event with time interval.
	k.Every(func() {
		fmt.Printf("%d times\n", count)
		count++
	})
}
