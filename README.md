# Krone

**Krone** is a minimalist package for creating scheduled tasks, which are triggered with timeout or at time intervals, similar to JavaScript's [setTimeout](https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/setTimeout) and [setInterval](https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/setInterval).

![Go](https://img.shields.io/badge/Golang-1.14-blue.svg?logo=go&longCache=true&style=flat)

This package exclusively uses the standard library.

Internally, the time and context packages are used, in addition to the channels to handle the events, the goroutines are not used internally, they are under the control of the client. So if you want the task to run in a separate goroutine, you must explicitly use the **go** keyword before executing the corresponding method, otherwise the program will stop until the event is executed.

## Getting Started

[Go](https://golang.org/) is required in version 1.7 or higher.

### Install

`go get -u github.com/orlmonteverde/krone`

### Features

* **Lightweight**, less than 100 lines of code.
* **Easy** to use.
* 100% compatible with the [time](https://godoc.org/time) and [context](https://godoc.org/context) packages of the **standard library**.
* It has **no external dependencies**.

## Examples

See [_examples/](https://github.com/orlmonteverde/krone/blob/master/_examples/) for a variety of examples.


### Timeout


```go
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
}

```

### Time interval


```go
package main

import (
	"fmt"
	"time"

	"github.com/orlmonteverde/krone"
)

func main() {
	// new krone instance.
	k := krone.New(2 * time.Second)

	// event with time interval.
	// Ctrl + C for stop.
	var count int
	k.Every(func() {
		fmt.Printf("%d times\n", count)
		count++
	})
}

```

### From time


```go
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


```

### With goroutines


```go
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
```

### With context


```go
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
```

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/orlmonteverde/krone/tags).

## Authors

* **Orlando Monteverde** - *Initial work* - [orlmonteverde](https://github.com/orlmonteverde)
