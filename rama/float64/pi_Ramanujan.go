package main

import (
	"fmt"
	"math"
)

func factorial(n float64) float64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var qPi float64
	var n float64
	for n = 0; ; n++ {
		den := math.Pow(-1, n) * factorial(4*n) * (1123 + 21460*n)
		num := math.Pow(882, 2*n+1) * math.Pow(math.Pow(4, n)*factorial(n), 4)
		ram := den / num
		qPi = qPi + ram
		//fmt.Printf("%v:\t%v / %v -> %v\t%v\n", n, den, num, ram, 4/qPi)
		fmt.Printf("%v:\t%v\n", n, 4/qPi)
	}
}
