package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/orlmonteverde/krone"
)

func main() {
	var wg sync.WaitGroup

	// new krone instance.
	k := krone.New(2 * time.Second)

	wg.Add(5)
	now := time.Now()

	// event with time interval.
	go k.Every(func() {
		fmt.Printf("In %v\n", time.Since(now))
		wg.Done()
	})

	// Wait until the event has fired 5 times.
	wg.Wait()
}
