package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var i, o float64
	for {
		rx := rand.Float64()
		ry := rand.Float64()
		if rx*rx+ry*ry <= 1 {
			i++
		} else {
			o++
		}
		fmt.Printf("%v:\t%v\n", i+o, 4*i/(i+o))
	}
}
