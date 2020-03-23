package main

import (
	"testing"
)

func benchmarkRamanujanPi(prec uint, b *testing.B) {
	for i := 0; i < b.N; i++ {
		pi(prec)
	}
}

/*
func BenchmarkRamanujanPi1(b *testing.B) {
	benchmarkRamanujanPi(1, b)
}

func BenchmarkRamanujanPi10(b *testing.B) {
	benchmarkRamanujanPi(10, b)
}

func BenchmarkRamanujanPi100(b *testing.B) {
	benchmarkRamanujanPi(100, b)
}

func BenchmarkRamanujanPi1000(b *testing.B) {
	benchmarkRamanujanPi(1000, b)
}
*/

func BenchmarkRamanujanPi10000(b *testing.B) {
	benchmarkRamanujanPi(10000, b)
}

func BenchmarkRamanujanPi20000(b *testing.B) {
	benchmarkRamanujanPi(20000, b)
}

func BenchmarkRamanujanPi30000(b *testing.B) {
	benchmarkRamanujanPi(30000, b)
}

func BenchmarkRamanujanPi40000(b *testing.B) {
	benchmarkRamanujanPi(40000, b)
}
