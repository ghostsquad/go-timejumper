package wrappers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func BenchmarkTime_Now(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

func BenchmarkJumperClock_Now(b *testing.B) {
	c := New()

	for i := 0; i < b.N; i++ {
		c.Now()
	}
}

func TestJumperClock_Now(t *testing.T) {
	t.Parallel()
	sleepTime := 50 * time.Millisecond
	fudgeTime := 10 * time.Millisecond

	for i := 0; i <= 100; i++ {
		t.Run("default behavior", func(t *testing.T) {
			c := New()

			realNow := time.Now()
			time.Sleep(sleepTime)
			cNow := c.Now()

			diffNs := cNow.Sub(realNow).Nanoseconds()

			//fmt.Printf("diff: %+v\n", diff)

			// To avoid race conditions here, I'm allowing a small difference in actual timing
			// This might require more research
			// The benchmarks above prove that .Now() is actually very fast
			assert.GreaterOrEqual(t, diffNs, sleepTime.Nanoseconds())
			assert.LessOrEqual(t, diffNs, (sleepTime + fudgeTime).Nanoseconds())
		})
	}

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
