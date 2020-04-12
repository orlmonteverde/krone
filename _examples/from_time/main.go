package main

import (
	"fmt"
	"time"

	"github.com/orlmonteverde/krone"
)

func main() {
	now := time.Now()
	date := now.Add(10 * time.Second)

	// new krone instance from time.
	k := krone.FromTime(date)

	k.Do(func() {
		fmt.Printf("ends in %v\n", time.Since(now))
	})
}
