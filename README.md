# TimeJumper

[![Documentation](https://godoc.org/github.com/ghostsquad/go-timejumper?status.svg)](http://godoc.org/github.com/ghostsquad/go-timejumper) [![Go Report Card](https://goreportcard.com/badge/github.com/ghostsquad/go-timejumper)](https://goreportcard.com/report/github.com/ghostsquad/go-timejumper)

Time as a dependency, very useful for testing. Unlike other similar libraries that use package globals, this gives you `Clock` interface that you can use in structs and functions as a dependency. `RealClock{}` delegates to the `time` standard library package.

## Quick Start

```go
package main

import (
    "time"
    timejumper "github.com/ghostsquad/go-timejumper"
)

clock := timejumper.New()
t := clock.Now()
fmt.Printf("current: %v\n", t)

fmt.Println("Dive into the future!")
clock.Jump(t.AddDate(1, 0, 0))
fmt.Printf("future: %v\n", clock.Now())

fmt.Println("Sleep, but not really")
clock.Sleep(10 * time.Hour)
fmt.Printf("future: %v\n", clock.Now())

fmt.Println("Back to the present")
clock.Back()
fmt.Printf("current: %v\n", clock.Now())

fmt.Println("Back to the past")
clock.Freeze(t)
fmt.Printf("frozen current: %v\n", clock.Now())
time.Sleep(5 * time.Seconds)
fmt.Printf("frozen current: %v\n", clock.Now())
```

## Development

### Testing

```bash
go test -bench=.
go test .
```
