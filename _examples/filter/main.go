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

	// The event will only be executed for cases in which the
	// function returns true.
	k.FilterFunc = func(t time.Time) bool {
		return t.After(now.Add(10 * time.Second))
	}

	// event with time interval.
	// Ctrl + C for stop.
	k.Every(func() {
		fmt.Printf("In %v\n", time.Since(now))
	})
}
