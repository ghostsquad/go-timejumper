package timejumper

import (
	"testing"
	"time"
)

func TestJumperClock_Race(t *testing.T) {
	c := New()

	go func() {
		for {
			c.Freeze(time.Now())
			c.Now()
		}
	}()

	go func() {
		for {
			c.Jump(time.Now())
			c.Now()
		}
	}()

	go func() {
		for {
			c.Sleep(10 * time.Millisecond)
			c.Now()
		}
	}()

	go func() {
		for {
			c.Back()
			c.Now()
		}
	}()

	go func() {
		for {
			now := c.Now()
			c.Since(now)
		}
	}()

	time.Sleep(1 * time.Second)
}
