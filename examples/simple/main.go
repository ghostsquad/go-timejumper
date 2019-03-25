package main

import (
	"fmt"
	"time"

	timejumper "github.com/ghostsquad/go-timejumper"
)

func main() {
	clock := timejumper.New()
	t := clock.Now()

	fmt.Printf("current: %v\n", t)

	fmt.Println("")

	fmt.Println("Dive into the future!")
	clock.Jump(t.AddDate(0, 0, 11))
	fmt.Printf("future: %v\n", clock.Now())

	fmt.Println("")

	fmt.Println("Sleep, but not really")
	clock.Sleep(10 * time.Second)
	fmt.Printf("future: %v\n", clock.Now())

	fmt.Println("")

	fmt.Println("Back to the present")
	clock.Back()
	fmt.Printf("current: %v\n", clock.Now())

	fmt.Println("")

	fmt.Println("Freeze the present")
	clock.Freeze(t)
	fmt.Printf("frozen current: %v\n", clock.Now())
	fmt.Println("Doing an actual 2 sec sleep for demonstration...")
	time.Sleep(2 * time.Second)
	fmt.Printf("frozen current: %v\n", clock.Now())
}
