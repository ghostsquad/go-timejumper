// +build !race

package timejumper

import (
	"testing"
	"time"
)

func BenchmarkTime_Now(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

func BenchmarkRealClock_Now(b *testing.B) {
	c := RealClock{}

	for i := 0; i < b.N; i++ {
		c.Now()
	}
}

func BenchmarkJumperClock_Now(b *testing.B) {
	c := New()

	for i := 0; i < b.N; i++ {
		c.Now()
	}
}
