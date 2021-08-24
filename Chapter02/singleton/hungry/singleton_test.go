package singleton

import (
	"testing"

	"github.com/chyidl/Chapter02/singleton/hungry/singleton"
	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	// assert equality
	assert.Equal(t, singleton.GetInstance(), singleton.GetInstance(), "they should be equal")
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetGetInstance() != singleton.GetInstance {
				b.Errorf("test fail")
			}
		}
	})
}
