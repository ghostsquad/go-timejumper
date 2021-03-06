package timejumper

// https://stackoverflow.com/questions/18970265/is-there-an-easy-way-to-stub-out-time-now-globally-in-golang-during-test

import (
	"errors"
	"sync"
	"time"
)

// Clock is an interface that allows you to a create
// different implementions for testing purposes
type Clock interface {
	Now() time.Time
	Since(time.Time) time.Duration
	Sleep(time.Duration)
}

// RealClock implements the standard lib
// only functions relevant for this codebase
type RealClock struct{}

// Now is a proxy to time.Now
func (RealClock) Now() time.Time {
	return time.Now()
}

// Since is a proxy to time.Since
func (RealClock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

// Sleep is a proxy to time.Sleep
func (RealClock) Sleep(d time.Duration) {
	time.Sleep(d)
}

// compile-time assertion that RealClock matches Clock interface
var _ Clock = RealClock{}

// TimeMachine is a function that gets called each time Now() is called
// The default machine allows the usage of freeze, scale and jump features of
// Jumper. This allows you to completely switch out the functionality
type TimeMachine func() (newTime time.Time)

// JumperClock is a Go implementation of but does not use package globals
// https://github.com/travisjeffery/timecop
// Instantiate with New()
type JumperClock struct {
	isFrozen         bool
	initialTime      time.Time
	initialTimeSetAt time.Time
	scale            int
	activate         TimeMachine
	mutex            sync.Mutex
}

// New returns a new instance of a TimeCopClock
func New() *JumperClock {
	c := &JumperClock{
		scale: 1,
	}

	c.Back()
	c.activate = c.NewDefaultTimeMachine()

	return c
}

// Now activates the TimeMachine
func (c *JumperClock) Now() time.Time {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.activate()
}

// Since returns the time elapsed since t. It is shorthand for .Now().Sub(t).
func (c *JumperClock) Since(t time.Time) time.Duration {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.activate().Sub(t)
}

// Sleep is similar to jump, except that it takes a duration
func (c *JumperClock) Sleep(d time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := c.activate()

	// Sleep for the scaled duration
	c.initialTime = now.Add(d * time.Duration(c.scale))
	c.initialTimeSetAt = time.Now()
}

// compile-time assertion that RealClock matches Clock interface
var _ Clock = &JumperClock{}

// NewDefaultTimeMachine is what's used to enable scaling and freezing
func (c *JumperClock) NewDefaultTimeMachine() TimeMachine {
	return func() time.Time {
		if c.isFrozen {
			return c.initialTime
		}

		diff := time.Now().Sub(c.initialTimeSetAt)
		return c.initialTime.Add(diff * time.Duration(c.scale))
	}
}

// Freeze ensures time doesn't progress forward
// You can still move time by freezing again with a new time
// or by using Travel
func (c *JumperClock) Freeze(t time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.isFrozen = true
	c.initialTime = t
	c.initialTimeSetAt = time.Now()
}

// Jump sets the clock to the desired time
// Compatible with Freeze (will remain frozen if previously frozen)
func (c *JumperClock) Jump(t time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.initialTime = t
	c.initialTimeSetAt = time.Now()
}

// Scale sets the time scale multiplier
// If set to 0, freezes the clock instead
func (c *JumperClock) Scale(s int) error {
	if s < 0 {
		return errors.New("Cannot set scale to less than 0")
	}

	if s == 0 {
		c.isFrozen = true
		return nil
	}

	c.scale = s
	return nil
}

// Back returns you to the present (future? since you started)
// Does not change the TimeMachine
func (c *JumperClock) Back() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()

	c.isFrozen = false
	c.initialTime = now
	c.initialTimeSetAt = now
}

// SetTimeMachine sets TimeMachine, a function that controls how time works!
// It also resets the clock according to Back()
func (c *JumperClock) SetTimeMachine(t TimeMachine) {
	c.Back()
	c.activate = t
}
