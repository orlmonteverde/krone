package main

import (
	"fmt"
	"time"

	"github.com/orlmonteverde/krone"
)

func main() {
	// new krone instance.
	k := krone.New(2 * time.Second)

	now := time.Now()

	// event with timeout.
	k.Do(func() {
		fmt.Printf("ends after %v \n", time.Since(now))
	})

	// event with time interval.
	// Ctrl + C for stop.
	var count int
	k.Every(func() {
		fmt.Printf("%d times\n", count)
		count++
	})
}
