// +build !race

package timejumper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJumperClock_Now(t *testing.T) {
	t.Parallel()
	sleepTime := 10 * time.Millisecond

	t.Run("frozen behavior", func(t *testing.T) {
		c := New()

		present := time.Now()
		c.Freeze(present)
		time.Sleep(sleepTime)
		future := c.Now()

		assert.Equal(t, present, future)
	})

	t.Run("scaling behavior", func(t *testing.T) {
		c := New()
		
		present := time.Now()
		c.Freeze(present)
		scale := 2
		c.Scale(scale)
		c.Sleep(sleepTime)
		future := c.Now()

		expectedDiff := sleepTime * time.Duration(scale)

		assert.Equal(t, expectedDiff, future.Sub(present))	
	})

	t.Run("jumping behavior", func(t *testing.T) {
		c := New()
		
		present := time.Now()
		c.Freeze(present)
		future := present.AddDate(1, 0, 0)
		c.Jump(future)

		assert.Equal(t, future, c.Now())
	})

	t.Run("sleeping behavior", func(t *testing.T) {
		c := New()
		present := time.Now()
		c.Freeze(present)
		c.Sleep(sleepTime)

		assert.Equal(t, present.Add(sleepTime), c.Now())
	})
}
