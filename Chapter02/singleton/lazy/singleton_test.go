package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	// assert equality
	assert.Equal(t, GetInstance(), GetInstance(), "they should be equal")
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

/*
go-design-patterns/Chapter02/singleton/lazy on  main [!] via 🐹 v1.16.6
➜ go test -bench=. -v
=== RUN   TestGetInstance
--- PASS: TestGetInstance (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/chyidl/go-design-patterns/Chapter02/singleton/lazy
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkGetInstanceParallel
BenchmarkGetInstanceParallel-8   	1000000000	         0.7532 ns/op
PASS
ok  	github.com/chyidl/go-design-patterns/Chapter02/singleton/lazy	1.010s
*/
