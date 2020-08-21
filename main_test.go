package main

import (
	"testing"
)

// go test -bench . -benchmem
func BenchmarkDoSmth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoSmth()
	}
}
