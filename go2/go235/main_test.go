package main

import (
	"testing"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Original
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Ex 2.3
func TableLoop(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>uint(i*8))])
	}
	return sum
}

// Ex 2.4
func PopCountShiftValue(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

// Ex 2.5
func PopCountDiscardBit(x uint64) (value int) {
	for x != 0 {
		x &= x - 1
		value++
	}
	return value
}

// Ex 2.3
func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

// Ex 2.3
func BenchmarkTable(b *testing.B) {
	bench(b, PopCount)
}

// Ex 2.3
func BenchmarkTableLoop(b *testing.B) {
	bench(b, TableLoop)
}

// Ex 2.4
func BenchmarkTableShiftValue(b *testing.B) {
	bench(b, PopCountShiftValue)
}

// Ex 2.5
func BenchmarkTableDiscardBit(b *testing.B) {
	bench(b, PopCountDiscardBit)

}
