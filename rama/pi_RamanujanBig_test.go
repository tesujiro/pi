package main

import (
	"testing"
)

func BenchmarkRamanujanPi(b *testing.B) {
	var i uint
	for i = 1; i < uint(b.N); i++ {
		pi(i)
	}
}
