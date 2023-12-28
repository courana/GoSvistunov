package main

import (
	"testing"
)

func benchmarkRender(b *testing.B, worker int) {
	for i := 0; i < b.N; i++ {
		render(worker)
	}
}

func BenchmarkRender1(b *testing.B) {
	benchmarkRender(b, 1)
}

func BenchmarkRender2(b *testing.B) {
	benchmarkRender(b, 2)
}

func BenchmarkRender4(b *testing.B) {
	benchmarkRender(b, 4)
}

func BenchmarkRender8(b *testing.B) {
	benchmarkRender(b, 8)
}

func BenchmarkRender16(b *testing.B) {
	benchmarkRender(b, 16)
}

func BenchmarkRender32(b *testing.B) {
	benchmarkRender(b, 32)
}

func BenchmarkRender64(b *testing.B) {
	benchmarkRender(b, 64)
}
