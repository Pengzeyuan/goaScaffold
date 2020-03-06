package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandIntN(t *testing.T) {
	for i := 0; i < 10; i++ {
		num := RandIntnV2(100000)
		t.Logf("num: %d", num)
		assert.NotEqual(t, num, 0)
	}
}

func BenchmarkRandIntN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandIntnV2(100000)
	}
}
