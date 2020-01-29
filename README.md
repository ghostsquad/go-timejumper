# TimeJumper

[![Documentation](https://godoc.org/github.com/ghostsquad/go-timejumper?status.svg)](http://godoc.org/github.com/ghostsquad/go-timejumper) [![Go Report Card](https://goreportcard.com/badge/github.com/ghostsquad/go-timejumper)](https://goreportcard.com/report/github.com/ghostsquad/go-timejumper) [![Build Status](https://travis-ci.org/ghostsquad/go-timejumper.svg?branch=master)](https://travis-ci.org/ghostsquad/go-timejumper)

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
make install
make test
make test-bench
make test-race
```

### CI Testing

```bash
make test-ci
make test-ci-bench
make test-ci-race
```

Output files are contained in [reports/](reports/)

## Contribution

Thank you for your PR's! ❤️

Contribution, in any kind of way, is highly welcome!
It doesn't matter if you are not able to write code.
Creating issues or holding talks and help other people to use [go-timejumper](https://github.com/ghostsquad/go-timejumper) is contribution, too!
A few examples:

* Correct typos in the README / documentation
* Reporting bugs
* Implement a new feature or endpoint
* Sharing the love of [go-timejumper](https://github.com/ghostsquad/go-timejumper) and help people to get use to it

If you are new to pull requests, checkout [Collaborating on projects using issues and pull requests / Creating a pull request](https://help.github.com/articles/creating-a-pull-request/) as well as our [pull request template](PULL_REQUEST_TEMPLATE.md)

## Releasing

Install [standard-version](https://github.com/conventional-changelog/standard-version)
```bash
npm i -g standard-version
```

```bash
standard-version
git push --tags
```

Manually copy/paste text from changelog (for this new version) into the release on Github.com. E.g.

[https://github.com/ghostsquad/go-timejumper/releases/edit/v0.1.1](https://github.com/ghostsquad/go-timejumper/releases/edit/v0.1.1)
