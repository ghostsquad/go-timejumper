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

		realNow := time.Now()
		c.Freeze(realNow)
		time.Sleep(sleepTime)
		cNow := c.Now()

		assert.Equal(t, realNow, cNow)
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

		assert.Equal(t, future - present, expectedDiff)	
	})

	t.Run("jumping behavior", func(t *testing.T) {
		c := New()
		
		present := time.Now()
		c.Freeze(present)
		future := realNow.AddDate(1, 0, 0)
		c.Jump(future)

		assert.Equal(t, c.Now(), future)
	})

	t.Run("sleeping behavior", func(t *testing.T) {
		c := New()
		realNow := time.Now()
		c.Freeze(realNow)
		c.Sleep(sleepTime)

		assert.Equal(t, realNow.Add(sleepTime), c.Now())
	})
}
