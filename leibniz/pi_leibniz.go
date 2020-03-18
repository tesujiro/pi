package main

import "fmt"

func main() {
	var qPi float64
	var sgn float64 = -1
	for n := 0; ; n++ {
		sgn = -1 * sgn
		qPi = qPi + sgn/float64(2*n+1)
		fmt.Printf("%v:\t%v\n", n, qPi*4)
	}
}
