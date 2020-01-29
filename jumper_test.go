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
	fudgeTime := 10 * time.Millisecond

	t.Run("frozen behavior", func(t *testing.T) {
		c := New()

		realNow := time.Now()
		c.Freeze(realNow)
		time.Sleep(sleepTime)
		cNow := c.Now()

		assert.Equal(t, realNow, cNow)
	})

	t.Run("frozen jumping behavior", func(t *testing.T) {
		c := New()

		realNow := time.Now()
		c.Freeze(realNow)

		future := realNow.AddDate(1, 0, 0)
		c.Jump(future)

		assert.Equal(t, future, c.Now())
	})

	t.Run("scaling behavior", func(t *testing.T) {
		c := New()

		scale := 2
		c.Scale(scale)

		realNow := time.Now()
		time.Sleep(sleepTime)
		cNow := c.Now()

		diffNs := cNow.Sub(realNow).Nanoseconds()
		expectedDiffMin := sleepTime * time.Duration(scale)
		expectedDiffMax := expectedDiffMin + fudgeTime

		assert.GreaterOrEqual(t, diffNs, expectedDiffMin.Nanoseconds())
		assert.LessOrEqual(t, diffNs, expectedDiffMax.Nanoseconds())
	})

	t.Run("jumping behavior", func(t *testing.T) {
		c := New()
		realNow := time.Now()
		future := realNow.AddDate(1, 0, 0)

		c.Jump(future)
		time.Sleep(sleepTime)
		cNow := c.Now()

		diffNs := cNow.Sub(realNow).Nanoseconds()
		expectedDiffMin := future.Sub(realNow) + sleepTime
		expectedDiffMax := expectedDiffMin + fudgeTime

		assert.GreaterOrEqual(t, diffNs, expectedDiffMin.Nanoseconds())
		assert.LessOrEqual(t, diffNs, expectedDiffMax.Nanoseconds())
	})

	t.Run("sleeping behavior", func(t *testing.T) {
		c := New()
		realNow := time.Now()
		c.Freeze(realNow)
		c.Sleep(sleepTime)
		cNow := c.Now()

		assert.Equal(t, realNow.Add(sleepTime), cNow)
	})
}
